package auth

import (
	"errors"
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
			ExpiresAt: now.Add(time.Second * 60).Unix(),
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

func ParseToken(accessToken string) (uint, error) {
	token, err := jwt.ParseWithClaims(accessToken, &claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*claims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}