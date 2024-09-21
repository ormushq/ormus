package authservice

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

type Config struct {
	SecretKey                  string `koanf:"secret_key"`
	ContextKey                 string `koanf:"context_key"`
	AccessExpirationTimeInDay  int    `koanf:"access_expiration_time_in_day"`
	AccessExpirationTTL        time.Duration
	RefreshExpirationTimeInDay int `koanf:"refresh_expiration_time_in_day"`
	RefreshExpirationTTL       time.Duration
	AccessSubject              string `koanf:"access_subject"`
	RefreshSubject             string `koanf:"refresh_subject"`
}

type Service struct {
	configs Config
}

func New(cfg Config) Service {
	hoursInDay := 24

	cfg.AccessExpirationTTL = time.Duration(cfg.AccessExpirationTimeInDay * hoursInDay * int(time.Hour))
	cfg.RefreshExpirationTTL = time.Duration(cfg.RefreshExpirationTimeInDay * hoursInDay * int(time.Hour))

	return Service{
		configs: cfg,
	}
}

func (s Service) GetConfig() Config {
	return s.configs
}

func (s Service) CreateAccessToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.configs.AccessSubject, s.configs.AccessExpirationTTL)
}

func (s Service) CreateRefreshToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.configs.RefreshSubject, s.configs.RefreshExpirationTTL)
}

func (s Service) ParseToken(bearerToken string) (*Claims, error) {
	// https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-ParseWithClaims-CustomClaimsType

	tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.GetConfig().SecretKey), nil
	})
	if err != nil {
		return nil, richerror.New("parse token failed").WithWrappedError(err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, richerror.New("parse token failed")
}

func (s Service) createToken(userID, subject string, expireDuration time.Duration) (string, error) {
	if userID == "" {
		// it is weird to build a jwt token for no one, right?
		return "", richerror.New("jwt.createToken").WithMessage(errmsg.ErrJwtEmptyUser)
	}

	// set our claims
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
		},
		UserID: userID,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(s.configs.SecretKey))
	if err != nil {
		return "", richerror.New("jwt.createToken").WithWrappedError(err)
	}

	return tokenString, nil
}
