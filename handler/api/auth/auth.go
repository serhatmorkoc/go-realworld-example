package auth

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(id uint) (string, error) {
	now := time.Now()
	claims := &claims{
		UserID: id,
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: now.Add(time.Hour * 60).Unix(),
			Id:        "",
			IssuedAt:  0,
			Issuer:    "",
			NotBefore: 0,
			Subject:   "",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return t, nil
}