package authservice

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	jwt.RegisteredClaims
	EnUserEmail string `json:"user_email"`
}

// TODO : we need first select encryption library than impelemnt this funtions
// TODO : if we want use encrypt/decrypt other pakages; i think it's better impelemnt this funtions in pkg/cryption.
type ClaimsEncryption interface {
	Encrypt(plainData string) (string, error)
	Decrypt(encryptedData string) (string, error)
}
