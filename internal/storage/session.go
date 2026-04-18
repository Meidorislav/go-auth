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

func (db *Database) GetSessionByToken(ctx context.Context, tokenHash string) (*Session, error) {
	var session Session
	err := db.Pool.QueryRow(ctx,
		`SELECT id, user_id, family_id, token_hash, user_agent, ip, expires_at, created_at
		FROM sessions WHERE token_hash = $1`, tokenHash,
	).Scan(&session.ID, &session.UserID, &session.FamilyID, &session.TokenHash,
		&session.UserAgent, &session.IP, &session.ExpiresAt, &session.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (db *Database) DeleteSession(ctx context.Context, id uuid.UUID) error {
	_, err := db.Pool.Exec(ctx,
		"DELETE FROM sessions WHERE id = $1", id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) DeleteSessionsByFamilyID(ctx context.Context, familyID uuid.UUID) error {
	_, err := db.Pool.Exec(ctx,
		"DELETE FROM sessions WHERE id = &1", familyID,
	)
	if err != nil {
		return err
	}

	return nil
}
