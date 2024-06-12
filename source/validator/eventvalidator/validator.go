package eventvalidator

import (
	"fmt"
)

type Repository interface {
}

type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{repo: repo}
}

type ValidatorError struct {
	Fields map[string]string `json:"error"`
	Err    error             `json:"message"`
}

func (v ValidatorError) Error() string {
	var err string

	for key, value := range v.Fields {
		err += fmt.Sprintf("%s: %s\n", key, value)
	}

	return err
}
