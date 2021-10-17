package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"os"
	"time"
)

type jwtCustomClaims struct {
	UserId   uint   `json:"user_id"`
	UserName string `json:"user_name"`
	UID      string `json:"uid"`
	jwt.StandardClaims
}

func GenerateToken(id uint, userName string) (string, error) {
	now := time.Now()
	claims := &jwtCustomClaims{
		UserId: id,
		UserName: userName,
		UID:    uuid.New().String(),
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
	jwtSecret := []byte(os.Getenv("JWT_ACCESS_SECRET"))
	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return t, nil
}

func ParseToken(accessToken string) (uint, error) {
	token, err := jwt.ParseWithClaims(accessToken, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("JWT_ACCESS_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*jwtCustomClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
