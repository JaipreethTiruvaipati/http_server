package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// HashPassword takes a plaintext string and scrambles it into an Argon2 hash
func HashPassword(password string) (string, error) {
	// CreateHash uses balanced defaults under the hood for you
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hash, nil
}

// CheckPasswordHash takes a plaintext guess, securely overlaps it against the database hash, and tells us if they match
func CheckPasswordHash(password, hash string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false, err
	}
	return match, nil
}

// MakeJWT generates an encrypted JWT string referencing the user's UUID
func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	signingKey := []byte(tokenSecret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "chirpy-access",
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
		Subject:   userID.String(),
	})
	return token.SignedString(signingKey)
}

// ValidateJWT verifies the signature on an incoming token and extracts the assigned User UUID
func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	// Initialize an empty struct to unpack the claims into
	claimsStruct := jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) {
			// This tells the JWT library what the decoding key is
			return []byte(tokenSecret), nil
		},
	)
	// Will fire if it's expired, corrupted, or used the wrong secret key
	if err != nil {
		return uuid.Nil, err
	}
	// Pull the User ID string out of the "Subject" field
	userIDString, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, err
	}
	// Double check the issuer is strictly our application
	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return uuid.Nil, err
	}
	if issuer != "chirpy-access" {
		return uuid.Nil, errors.New("invalid issuer")
	}
	// Convert the raw string safely back into a UUID
	id, err := uuid.Parse(userIDString)
	if err != nil {
		return uuid.Nil, errors.New("invalid user ID")
	}
	return id, nil
}

// GetBearerToken extracts the internal JWT from the standard Authorization header
func GetBearerToken(headers http.Header) (string, error) {
	// Pull the exact header out using Go's built in map
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no authorization header included")
	}
	// The frontend will always prefix the token with "Bearer ", so we look for it
	prefix := "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		return "", errors.New("malformed authorization header")
	}
	// Return exactly what is left after we chop the prefix off!
	return strings.TrimPrefix(authHeader, prefix), nil
}

// MakeRefreshToken generates a totally random 256-bit hex string perfectly suited for long-lived tracking
func MakeRefreshToken() (string, error) {
	b := make([]byte, 32) // 32 bytes safely = 256 bits of data
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
