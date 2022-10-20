package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jellydator/ttlcache/v3"
	"log"
	"os"
	"time"
)

type jwtService struct {
	cache *ttlcache.Cache[string, string]
}

func NewJwtService() *jwtService {
	cache := ttlcache.New[string, string](
		ttlcache.WithTTL[string, string](30 * time.Minute),
	)

	go cache.Start()

	return &jwtService{
		cache: cache,
	}
}

func (s *jwtService) Sign(payload map[string]interface{}, duration time.Duration) (string, error) {
	key := os.Getenv("APP_KEY")

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(duration * time.Minute).Unix()

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

func (s *jwtService) Verify(tokenString string) (*jwt.Token, error) {
	key := os.Getenv("APP_KEY")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// check blacklist
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	if blacklist := s.cache.Get(id); blacklist != nil {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func (s *jwtService) Revoke(tokenString string) error {
	key := os.Getenv("APP_KEY")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		log.Println(err.Error())
		return err
	}

	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	expF64 := claims["exp"].(float64)
	expTtl := time.Until(time.Unix(int64(expF64), 0))

	// blacklist
	s.cache.Set(id, tokenString, expTtl)

	return nil
}
