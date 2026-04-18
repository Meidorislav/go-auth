package storage

import (
	"context"
	"time"

	"github.com/gofrs/uuid/v5"
)

func (db *Database) CreateSession(ctx context.Context, userID, familyID uuid.UUID,
	tokenHash, userAgent, ip string, expiresAt time.Time) error {
	_, err := db.Pool.Exec(ctx,
		`INSERT INTO sessions (user_id, family_id, token_hash, user_agent, ip, expires_at) 
		VALUES ($1, $2, $3, $4, $5, $6)`,
		userID, familyID, tokenHash, userAgent, ip, expiresAt,
	)
	if err != nil {
		return err
	}

	return nil
}
