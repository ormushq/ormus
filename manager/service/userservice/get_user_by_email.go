package userservice

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s Service) GetUserByEmail(email string) (*entity.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, richerror.New("GetUserByEmail").WhitWarpError(err)
	}

	return user, nil
}
