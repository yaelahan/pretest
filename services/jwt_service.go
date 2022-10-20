package services

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtService interface {
	Sign(map[string]interface{}, time.Duration) (string, error)
	Verify(string) (*jwt.Token, error)
	Revoke(string) error
}
