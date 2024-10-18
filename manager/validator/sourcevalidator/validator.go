package sourcevalidator

import (
	"errors"
	"fmt"
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

type Repository interface {
	Exist(id string) (bool, error)
}

type Validator struct {
	sourceRepo Repository
}

func New(sourceRepo Repository) Validator {
	return Validator{
		sourceRepo: sourceRepo,
	}
}

func (v Validator) isSourceExist(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("error while reflection interface")
	}

	exist, err := v.sourceRepo.Exist(s)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("source does not exist")
	}

	return nil
}
