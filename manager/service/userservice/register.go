package userservice

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/managerparam"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/password"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s Service) Register(req param.RegisterRequest) (*param.RegisterResponse, error) {
	vErr := s.userValidator.ValidateRegisterRequest(req)
	if vErr != nil {
		return nil, vErr
	}
	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		return nil, richerror.New("register.hash").WithWrappedError(err)
	}

	user := entity.User{
		DeletedAt: nil,
		Email:     req.Email,
		Password:  hashedPassword,
		// Now it's true by default until our authentication system works properly.
		// TODO: The user is set to active when he has confirmed his email
		IsActive: true,
	}
	createdUser, err := s.repo.Register(user)
	if err != nil {
		return nil, richerror.New("register.repo").WithWrappedError(err)
	}

	// TODO: we have to trigger an event of registration in this phase of function
	// return create new user
	inOutChan, err := s.internalBroker.GetInputChannel(managerparam.CreateDefaultProject)
	if err != nil {
		return nil, richerror.New("register.internalBroker").WithWrappedError(err)
	}
	inOutChan <- []byte(createdUser.ID)

	return &param.RegisterResponse{
		ID:    createdUser.ID,
		Email: createdUser.Email,
	}, nil
}
