package userservice

import "github.com/ormushq/ormus/manager/entity"

type Repository interface {
	Register(u entity.User) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	DoesUserExistsByEmail(email string) (bool, error)
}

type JWTEngine interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}

type Service struct {
	repo Repository
	jwt  JWTEngine
}

// This benchmark is the result of using a pointer, or a struct in return of New() function of this package
//
//goos: darwin
//goarch: arm64
//pkg: github.com/ormushq/ormus/manager/service
//BenchmarkServiceStructReturn-10         14631142                78.64 ns/op          240 B/op          3 allocs/op
//BenchmarkServicePointerReturn-10        15361818                78.75 ns/op          240 B/op          3 allocs/op
//PASS
//ok      github.com/ormushq/ormus/manager/service        2.590s

func New(authGenerator JWTEngine, repository Repository) *Service {
	return &Service{jwt: authGenerator, repo: repository}
}
