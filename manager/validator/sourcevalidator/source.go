package sourcevalidator

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/oklog/ulid/v2"
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/managerparam"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (v Validator) ValidateCreateSourceForm(req managerparam.AddSourceRequest) *ValidatorError {
	minNameLength := 5
	maxNameLength := 30

	minDescriptionLength := 5
	maxDescriptionLength := 100

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(minNameLength, maxNameLength), validation.By(v.isSourceAlreadyCreated)),
		validation.Field(&req.Description, validation.Required, validation.Length(minDescriptionLength, maxDescriptionLength)),
		validation.Field(&req.ProjectID, validation.Required, validation.By(v.validateULID)),
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
			Err: richerror.New("sourcevalidator.ValidateCreateSourceForm").WithMessage(errmsg.ErrorMsgInvalidInput).WhitKind(richerror.KindInvalid).
				WhitMeta(map[string]interface{}{"request:": req}).WithWrappedError(err),
		}
	}

	return nil
}

func (v Validator) ValidateUpdateSourceForm(req managerparam.UpdateSourceRequest) *ValidatorError {
	minNameLength := 5
	maxNameLength := 30

	minDescriptionLength := 5
	maxDescriptionLength := 100

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(minNameLength, maxNameLength)),
		validation.Field(&req.Description, validation.Required, validation.Length(minDescriptionLength, maxDescriptionLength)),
		validation.Field(&req.ProjectID, validation.Required, validation.By(v.validateULID)),
		validation.Field(&req.Status, validation.Required, validation.By(v.validateStatus)),
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
			Err: richerror.New("sourcevalidator.ValidateUpdateSourceForm").WithMessage(errmsg.ErrorMsgInvalidInput).WhitKind(richerror.KindInvalid).
				WhitMeta(map[string]interface{}{"request:": req}).WithWrappedError(err),
		}
	}

	return nil
}

func (v Validator) ValidateIDToDelete(id string) *ValidatorError {
	if err := validation.Validate(id,
		validation.By(v.isSourceAlreadyCreated),
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
			Err: richerror.New("sourcevalidator.ValidateIDToDelete").WithMessage(errmsg.ErrorMsgInvalidInput).WhitKind(richerror.KindInvalid).
				WhitMeta(map[string]interface{}{"request:": id}).WithWrappedError(err),
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
		return errors.New("invalid id")
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

func (v Validator) validateStatus(value interface{}) error {
	s, ok := value.(entity.Status)
	if !ok {
		return errors.New("error while reflection interface")
	}

	switch s {
	case "active":
		return nil
	case "not active":
		return nil
	}

	return errors.New("invalide status of source")
}
