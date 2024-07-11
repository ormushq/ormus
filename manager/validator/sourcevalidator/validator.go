package sourcevalidator

import (
	"fmt"

	"github.com/ormushq/ormus/manager/service/sourceservice"
)

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

type Validator struct {
	repo sourceservice.SourceRepo
}

func New(repo sourceservice.SourceRepo) *Validator {
	return &Validator{repo: repo}
}
