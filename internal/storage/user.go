package storage

import "context"

func (db *Database) CreateUser(ctx context.Context, email, password_hash string) error {
	_, err := db.Pool.Exec(ctx,
		"INSERT INTO users (email, password_hash) VALUES ($1, $2)",
		email, password_hash,
	)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := db.Pool.QueryRow(ctx,
		"SELECT id, email, password_hash, created_at, updated_at FROM users WHERE email = $1",
		email,
	).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
