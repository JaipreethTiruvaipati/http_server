package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestJWTGenerationAndValidation(t *testing.T) {
	secret := "super-secret-testing-key"
	userID := uuid.New()
	duration := time.Hour

	// 1. Happy Path: Make sure a standard token is created and decodes cleanly back to the same ID
	token, err := MakeJWT(userID, secret, duration)
	if err != nil {
		t.Fatalf("Failed to make JWT: %v", err)
	}

	validatedID, err := ValidateJWT(token, secret)
	if err != nil {
		t.Fatalf("Failed to validate JWT: %v", err)
	}

	if validatedID != userID {
		t.Errorf("Expected ID %v, got %v", userID, validatedID)
	}

	// 2. Reject mismatched secrets
	_, err = ValidateJWT(token, "wrong-secret")
	if err == nil {
		t.Error("Expected error with wrong secret, got none")
	}

	// 3. Reject expired tokens
	// By giving it a negative duration, the token is instantly generated already "dead"
	expiredToken, _ := MakeJWT(userID, secret, -time.Hour)
	_, err = ValidateJWT(expiredToken, secret)
	if err == nil {
		t.Error("Expected error with expired token, got none")
	}
}
