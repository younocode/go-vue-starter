package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/younocode/go-vue-starter/server/config"
	"time"
)

type JWT struct {
	secret   []byte
	duration time.Duration
}

func NewJWT(cfg config.JWTConfig) *JWT {
	return &JWT{
		secret:   []byte(cfg.Secret),
		duration: cfg.Duration,
	}
}

type UserClaims struct {
	Email  string `json:"emailSender"`
	UserId int32  `json:"user_id"`
	jwt.RegisteredClaims
}

func (j *JWT) Generate(email string, userID int32) (string, error) {
	claims := UserClaims{
		Email:  email,
		UserId: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}

func (j *JWT) ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func (j *JWT) GenerateRefreshToken(email string, userID int32) (string, error) {
	claims := UserClaims{
		Email:  email,
		UserId: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}

func (j *JWT) ParseRefreshToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
