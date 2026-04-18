package token

import (
	"crypto/rsa"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt"
)

func GenerateAccessToken(userID uuid.UUID, privateKey *rsa.PrivateKey) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID.String(),
		"exp": time.Now().Add(15 * time.Minute).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}
