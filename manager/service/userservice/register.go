package userservice

import (
	"encoding/json"
	"github.com/ormushq/ormus/pkg/channel"

	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/managerparam"
	"github.com/ormushq/ormus/manager/managerparam/projectparam"
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
	// TODO move this to init service and store input channel to service struct
	inputChan, err := s.internalBroker.GetInputChannel(managerparam.CreateDefaultProject)
	if err != nil {
		// TODO don`t send trigger error to client just log it
		return nil, richerror.New("register.internalBroker").WithWrappedError(err)
	}
	createProjectReq := projectparam.CreateThoughChannel{
		UserID:      createdUser.ID,
		Name:        "Default",
		Description: "Default project",
	}
	js, err := json.Marshal(createProjectReq)
	if err != nil {
		logger.L().Error(err.Error())
	} else {
		inputChan <- channel.Message{Body: js}
	}

	return &param.RegisterResponse{
		ID:    createdUser.ID,
		Email: createdUser.Email,
	}, nil
}
