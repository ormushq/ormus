package sourcevalidator

import (
	"errors"
	"fmt"
	"github.com/oklog/ulid/v2"

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

func New(repo sourceservice.SourceRepo) Validator {
	return Validator{repo: repo}
}

func (v Validator) validateULID(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("error while reflection interface")
	}

	_, err := ulid.Parse(s)
	if err != nil {
		return errors.New("invalid id")
	}

	return nil
}
