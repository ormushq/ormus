package uservalidator

import (
	"fmt"
	"github.com/ormushq/ormus/manager/service/userservice"
)

const (
	emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
)

type ValidatorError struct {
	Fields map[string]string
	Err    error
}

func (v ValidatorError) Error() string {
	var err string

	for key, value := range v.Fields {
		err += fmt.Sprintf("%s: %s\n", key, value)
	}

	return err
}

type Validator struct {
	repo userservice.Repository
}

func New(repo userservice.Repository) Validator {
	return Validator{repo: repo}
}
