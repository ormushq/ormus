package authservice

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/ormushq/ormus/manager/entity"
	"strings"
	"time"
)

const (
	defultSignKey               = "Ormus_jwt"
	defultAccessExpirationTime  = time.Hour * 24 * 7
	defultRefreshExpirationTime = time.Hour * 24 * 7 * 4
	defultAccessSubject         = "ac"
	defultRefreshSubject        = "rt"
)

type Config struct {
	signKey               string
	accessExpirationTime  time.Duration
	refreshExpirationTime time.Duration
	accessSubject         string
	refreshSubject        string
}

type Service struct {
	configs Config
}

var auth *Service

func init() {
	auth = &Service{
		configs: Config{
			signKey:               defultSignKey,
			accessExpirationTime:  defultAccessExpirationTime,
			refreshExpirationTime: defultRefreshExpirationTime,
			accessSubject:         defultAccessSubject,
			refreshSubject:        defultRefreshSubject,
		}}
}
func New(signKey, accessSubject, refreshSubject string,
	accessExpirationTime, refreshExpirationTime time.Duration) Service {
	return Service{
		configs: Config{
			signKey:               signKey,
			accessExpirationTime:  accessExpirationTime,
			refreshExpirationTime: refreshExpirationTime,
			accessSubject:         accessSubject,
			refreshSubject:        refreshSubject,
		}}
}
func (s Service) CreateAccessToken(user entity.User) (string, error) {
	return s.createToken(user.Email, s.configs.accessSubject, s.configs.accessExpirationTime)
}
func (s Service) CreateRefreshToken(user entity.User) (string, error) {
	return s.createToken(user.Email, s.configs.refreshSubject, s.configs.refreshExpirationTime)
}
func (s Service) ParseToken(bearerToken string) (*Claims, error) {
	//https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-ParseWithClaims-CustomClaimsType

	tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.configs.signKey), nil
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

func (s Service) createToken(userEmail, subject string, expireDuration time.Duration) (string, error) {
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
	tokenString, err := accessToken.SignedString([]byte(s.configs.signKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
