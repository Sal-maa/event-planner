package middleware

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct{}

var SECRET_KEY = []byte("iN1_4daL4H_53cR3tKey_m1LiK_juNd4n4")

func AuthService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userId int) (string, error) {
	payload := jwt.MapClaims{}
	payload["id"] = userId
	payload["expired"] = time.Now().Add(time.Hour * 5).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil
}
