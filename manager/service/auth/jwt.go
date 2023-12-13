package auth

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ormushq/ormus/manager/entity"
	errors "github.com/ormushq/ormus/manager/error"
)

var jwtModule *JWT

const (
	defaultSecretKey             = "Ormus_jwt"
	defaultAccessExpirationTime  = time.Hour * 24 * 7
	defaultRefreshExpirationTime = time.Hour * 24 * 7 * 4
	defaultAccessSubject         = "ac"
	defaultRefreshSubject        = "rt"
)

type JwtConfig struct {
	secretKey             string        `koanf:"secret_key"`
	accessExpirationTime  time.Duration `koanf:"access_expiration_time"`
	refreshExpirationTime time.Duration `koanf:"refresh_expiration_time"`
	accessSubject         string        `koanf:"access_subject"`
	refreshSubject        string        `koanf:"refresh_subject"`
}

type JWT struct {
	configs JwtConfig
}

func init() {
	jwtModule = &JWT{
		configs: JwtConfig{
			secretKey:             defaultSecretKey,
			accessExpirationTime:  defaultAccessExpirationTime,
			refreshExpirationTime: defaultRefreshExpirationTime,
			accessSubject:         defaultAccessSubject,
			refreshSubject:        defaultRefreshSubject,
		}}
}

func NewJWT(config *JwtConfig) *JWT {
	// if the config was not loaded return the jwt module with default config
	if config == nil {
		return jwtModule
	}

	return &JWT{
		configs: JwtConfig{
			secretKey:             config.secretKey,
			accessExpirationTime:  config.accessExpirationTime,
			refreshExpirationTime: config.refreshExpirationTime,
			accessSubject:         config.accessSubject,
			refreshSubject:        config.refreshSubject,
		}}
}

func (s JWT) CreateAccessToken(user entity.User) (string, error) {
	if len(user.Email) == 0 {
		// it is wierd to build a jwt token for no one, right?
		return "", errors.JwtEmptyUserErr
	}
	return s.createToken(user.Email, s.configs.accessSubject, s.configs.accessExpirationTime)
}

func (s JWT) CreateRefreshToken(user entity.User) (string, error) {
	return s.createToken(user.Email, s.configs.refreshSubject, s.configs.refreshExpirationTime)
}

func (s JWT) ParseToken(bearerToken string) (*Claims, error) {
	//https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-ParseWithClaims-CustomClaimsType

	tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.configs.secretKey), nil
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
	tokenString, err := accessToken.SignedString([]byte(s.configs.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
