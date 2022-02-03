package middleware

import (
	"context"
	"net/http"
	"strings"

	_models "github.com/justjundana/event-planner/models"

	jwt "github.com/golang-jwt/jwt"
)

var ctxKey = &contextKey{"user"}

type contextKey struct {
	data string
}

func Authentication() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			authHeader := req.Header.Get("Authorization")

			if !strings.Contains(authHeader, "Bearer") {
				next.ServeHTTP(res, req)
				return
			}

			tokenString := ""
			arrayToken := strings.Split(authHeader, " ")
			if len(arrayToken) == 2 {
				tokenString = arrayToken[1]
			}

			token, err := AuthService().ValidateToken(tokenString)
			if err != nil {
				http.Error(res, "invalid token", http.StatusForbidden)
				return
			}

			payload, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(res, "Unauthorized", http.StatusForbidden)
				return
			}

			userID := int(payload["id"].(float64))

			user := _models.User{ID: userID}

			ctx := context.WithValue(req.Context(), ctxKey, &user)

			req = req.WithContext(ctx)
			next.ServeHTTP(res, req)
		})
	}
}

func ForContext(ctx context.Context) *_models.User {
	raw, _ := ctx.Value(ctxKey).(*_models.User)
	return raw
}
