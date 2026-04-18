package main

import (
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/username/chirpy/internal/database"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	DB             *database.Queries
	Platform       string
	JWTSecret      string
	PolkaKey       string
}

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

type User struct {
	ID          uuid.UUID `json:"id"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsChirpyRed bool      `json:"is_chirpy_red"`
}
