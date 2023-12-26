package sourcevalidator

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/oklog/ulid/v2"
	"github.com/ormushq/ormus/manager/param"
)

func (v Validator) ValidateCreateSourceForm(req *param.AddSourceRequest) *ValidatorError {
	minNameLength := 5
	maxNameLength := 30

	minDescriptionLength := 5
	maxDescriptionLength := 100

	if err := validation.ValidateStruct(req,
		validation.Field(req.Name, validation.Required, validation.Length(minNameLength, maxNameLength), validation.By(v.isSourceAlreadyCreated)),
		validation.Field(req.Description, validation.Required, validation.Length(minDescriptionLength, maxDescriptionLength)),
		validation.Field(req.ProjectId, validation.Required, validation.By(v.validateULID)),
		validation.Field(req.OwnerId, validation.Required, validation.By(v.validateULID)),
	); err != nil {

		fieldErr := make(map[string]string)

		var errV validation.Errors
		ok := errors.As(err, &errV)

		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErr[key] = value.Error()
				}
			}
		}

		return &ValidatorError{
			Fields: fieldErr,
			Err:    errors.New("a"), // TODO wait for richerror
		}
	}

	return nil
}

func (v Validator) ValidateUpdateSourceForm(req *param.UpdateSourceRequest) *ValidatorError {

	minNameLength := 5
	maxNameLength := 30

	minDescriptionLength := 5
	maxDescriptionLength := 100

	if err := validation.ValidateStruct(req,
		validation.Field(req.Name, validation.Required, validation.Length(minNameLength, maxNameLength)),
		validation.Field(req.Description, validation.Required, validation.Length(minDescriptionLength, maxDescriptionLength)),
		validation.Field(req.ProjectId, validation.Required, validation.By(v.validateULID)),
	); err != nil {

		fieldErr := make(map[string]string)

		var errV validation.Errors
		ok := errors.As(err, &errV)

		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErr[key] = value.Error()
				}
			}
		}

		return &ValidatorError{
			Fields: fieldErr,
			Err:    errors.New("a"), // TODO wait for richerror
		}
	}

	return nil
}

func (v Validator) validateULID(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("error while reflection interface")
	}

	_, err := ulid.Parse(s)
	if err != nil {
		return err
	}

	return nil
}

func (v Validator) isSourceAlreadyCreated(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("error while reflection interface")
	}

	exist, err := v.repo.IsSourceAlreadyCreatedByName(s)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("this name is already usesd")
	}

	return nil
}
