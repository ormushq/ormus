package uservalidator

import (
	"github.com/ormushq/ormus/manager/service/userservice"
)

const (
	emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
)

type ValidatorError struct {
	Fields map[string]string
	Error  error
}

type Validator struct {
	repo userservice.Repository
}

func New(repo userservice.Repository) Validator {
	return Validator{repo: repo}
}
