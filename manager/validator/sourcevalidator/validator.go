package sourcevalidator

import (
	"errors"
	"fmt"
)

type ValidatorError struct {
	Fields map[string]string `json:"error"`
	Err    error             `json:"message"`
}
type Repo interface {
	IsExist(id string) (bool, error)
}

func (v ValidatorError) Error() string {
	var err string

	for key, value := range v.Fields {
		err += fmt.Sprintf("%s: %s\n", key, value)
	}

	return err
}

type Validator struct {
	sourceRepo  Repo
	projectRepo Repo
}

func New(sourceRepo, projectRepo Repo) Validator {
	return Validator{
		sourceRepo:  sourceRepo,
		projectRepo: projectRepo,
	}
}

func (v Validator) isSourceExist(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("error while reflection interface")
	}

	exist, err := v.sourceRepo.IsExist(s)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("source does not exist")
	}

	return nil
}

func (v Validator) isProjectExist(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("error while reflection interface")
	}

	exist, err := v.projectRepo.IsExist(s)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("project does not exist")
	}

	return nil
}
