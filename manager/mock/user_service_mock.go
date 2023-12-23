package usermock

import (
	"fmt"

	"github.com/ormushq/ormus/param"
)

const ServiceErr = "service error"

type MockService struct {
	serviceErr bool
	email      string
}

func NewMockService(serviceErr bool, email string) *MockService {
	return &MockService{serviceErr: serviceErr, email: email}
}

func (m MockService) Login(_ param.LoginRequest) (*param.LoginResponse, error) {
	if m.serviceErr {
		return nil, fmt.Errorf(ServiceErr)
	}

	return &param.LoginResponse{
		User: param.UserInfo{
			ID:    "new_id",
			Email: m.email,
		},
	}, nil
}

func (m MockService) Register(_ param.RegisterRequest) (*param.RegisterResponse, error) {
	if m.serviceErr {
		return nil, fmt.Errorf(ServiceErr)
	}

	return &param.RegisterResponse{
		Email: m.email,
	}, nil
}
