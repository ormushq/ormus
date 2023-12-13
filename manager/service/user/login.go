package userservice

import (
	"fmt"
	"github.com/ormushq/ormus/param"
)

func (s Service) Login(req param.LoginRequest) (param.LoginResponse, error) {
	//check the existing of email address  from repository
	//get user by email address
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return param.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	if user.Password != req.Password {
		return param.LoginResponse{}, fmt.Errorf("username or password isn't correct")
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
