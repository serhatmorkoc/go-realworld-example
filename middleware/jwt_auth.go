package middleware1

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/serhatmorkoc/go-realworld-example/handler/render"
	"net/http"
	"strings"
)

func ValidateJWT(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			render.Unauthorized(w, errors.New("authorization not found in header"))
			//http.Error(w, "authorization not found in header", http.StatusUnauthorized)
			return
		}
		jwtToken := authHeader[1]
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})
		if err != nil {
			render.Unauthorized(w, err)
			//http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "userAuthCtx", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		http.Error(w, "err.Error()", http.StatusUnauthorized)

	}
	return http.HandlerFunc(fn)
}
