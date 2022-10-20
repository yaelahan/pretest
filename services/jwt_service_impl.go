package services

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
	"time"
)

type jwtService struct{}

func NewJwtService() *jwtService {
	return &jwtService{}
}

func (*jwtService) Sign(payload map[string]interface{}, duration time.Duration) (string, error) {
	expiredAt := time.Now().Add(time.Minute * duration).Unix()

	key := os.Getenv("APP_KEY")

	claims := jwt.MapClaims{}
	claims["exp"] = expiredAt
	claims["authorization"] = true

	for i, v := range payload {
		claims[i] = v
	}

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := to.SignedString([]byte(key))

	if err != nil {
		log.Println(err.Error())
		return accessToken, err
	}

	return accessToken, nil
}

func (*jwtService) Verify(s string) (*jwt.Token, error) {
	key := os.Getenv("APP_KEY")
	token, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return token, nil
}
