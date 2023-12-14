package auth

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ormushq/ormus/manager/entity"
	errors "github.com/ormushq/ormus/manager/error"
)

type JwtConfig struct {
	SecretKey                  string        `koanf:"secret_key"`
	AccessExpirationTimeInDay  time.Duration `koanf:"access_expiration_time_in_day"`
	RefreshExpirationTimeInDay time.Duration `koanf:"refresh_expiration_time_in_day"`
	AccessSubject              string        `koanf:"access_subject"`
	RefreshSubject             string        `koanf:"refresh_subject"`
}

type JWT struct {
	configs JwtConfig
}

func NewJWT(cfg JwtConfig) *JWT {
	day := time.Hour * 24

	return &JWT{
		configs: JwtConfig{
			SecretKey:                  cfg.SecretKey,
			AccessExpirationTimeInDay:  cfg.AccessExpirationTimeInDay * day,
			RefreshExpirationTimeInDay: cfg.RefreshExpirationTimeInDay * day,
			AccessSubject:              cfg.AccessSubject,
			RefreshSubject:             cfg.RefreshSubject,
		},
	}
}

func (s JWT) CreateAccessToken(user entity.User) (string, error) {
	if len(user.Email) == 0 {
		// it is wierd to build a jwt token for no one, right?
		return "", errors.JwtEmptyUserErr
	}
	return s.createToken(user.Email, s.configs.AccessSubject, s.configs.AccessExpirationTimeInDay)
}

func (s JWT) CreateRefreshToken(user entity.User) (string, error) {
	return s.createToken(user.Email, s.configs.RefreshSubject, s.configs.RefreshExpirationTimeInDay)
}

func (s JWT) ParseToken(bearerToken string) (*Claims, error) {
	//https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-ParseWithClaims-CustomClaimsType

	tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.configs.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (s JWT) createToken(userEmail, subject string, expireDuration time.Duration) (string, error) {

	// create a signer for rsa 256
	// TODO - replace with rsa 256 RS256 - https://github.com/golang-jwt/jwt/blob/main/http_example_test.go

	// set our claims
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
		},
		UserEmail: userEmail,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(s.configs.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
