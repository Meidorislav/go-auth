package service

import (
	"context"

	"github.com/alexedwards/argon2id"
	"github.com/meidorislav/go-auth/internal/storage"
)

type Service struct {
	DB *storage.Database
}

func InitService(db *storage.Database) *Service {
	return &Service{
		DB: db,
	}
}

func HashPassword(password string) (string, error) {
	return argon2id.CreateHash(password, argon2id.DefaultParams)
}

func CheckingHash(password, hash string) (bool, error) {
	return argon2id.ComparePasswordAndHash(password, hash)
}

func (s *Service) Register(ctx context.Context, login, password string) error {
	hash, err := HashPassword(password)
	if err != nil {
		return err
	}
	s.DB.CreateUser(ctx, login, hash)
	return nil
}
