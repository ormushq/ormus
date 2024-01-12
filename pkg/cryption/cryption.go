package cryption

import "fmt"

type CryptConfing struct {
	CryptionSecretKey string `koanf:"cryption_secret_key"`
}

type Service struct {
	config CryptConfing
}

func New(config CryptConfing) *Service {
	return &Service{
		config: config,
	}
}

func (s Service) Encrypt(plainData string) (string, error) {
	fmt.Println(plainData, s.config.CryptionSecretKey)

	return "encryptUserEmail", nil
}

func (s Service) Decrypt(encryptedData string) (string, error) {
	fmt.Println(encryptedData, s.config.CryptionSecretKey)

	return "UserEmail", nil
}
