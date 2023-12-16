package userservice

import (
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/password"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s Service) Login(req param.LoginRequest) (param.LoginResponse, error) {
	// check the existing of email address  from repository
	// get user by email address
	user, err := s.repo.GetUserByEmail(req.Email)
	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		return param.LoginResponse{}, richerror.New("Login").WhitWarpError(err)
	}
	if user.Password != hashedPassword {
		return param.LoginResponse{}, richerror.New("Login").WhitWarpError(err).WhitMessage(errmsg.ErrWrongCredentials)
	}

	// jwt token
	AccessToken, err := s.jwt.CreateAccessToken(*user)
	if err != nil {
		return param.LoginResponse{}, richerror.New("Login").WhitWarpError(err)
	}
	RefreshToken, err := s.jwt.CreateRefreshToken(*user)
	if err != nil {
		return param.LoginResponse{}, richerror.New("Login").WhitWarpError(err)
	}

	// return ok
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
