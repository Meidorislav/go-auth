package storage

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type User struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Session struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	FamilyID  uuid.UUID
	TokenHash string
	UserAgent string
	IP        string
	ExpiresAt time.Time
	CreatedAt time.Time
}
