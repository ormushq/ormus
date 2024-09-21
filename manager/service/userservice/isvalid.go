package userservice

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s Service) IsUserIDValid(email string) (bool, error) {
	const op = "userservice.IsUserIDValid"

	user, rErr := s.repo.GetUserByEmail(email)
	if rErr != nil {
		return false, richerror.New(op).WithWrappedError(rErr).WithMessage(errmsg.ErrSomeThingWentWrong)
	}

	nilUser := entity.User{}
	if user == nilUser {
		return false, nil
	}

	return true, nil
}
