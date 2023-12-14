package auth

import (
	"fmt"
	ormuserror "github.com/ormushq/ormus/manager/error"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/password"
)

func (s Service) Login(req param.LoginRequest) (param.LoginResponse, error) {
	//check the existing of email address  from repository
	//get user by email address
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return param.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		return param.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	if user.Password != hashedPassword {
		return param.LoginResponse{}, ormuserror.ErrWrongCredentials
	}
	// jwt token
	AccessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return param.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	RefreshToken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		return param.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	//return ok
	return param.LoginResponse{
		User: param.UserInfo{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
			Email:     user.Email,
		},
		Tokens: param.Token{
			AccessToken:  AccessToken,
			RefreshToken: RefreshToken,
		},
	}, nil

}
