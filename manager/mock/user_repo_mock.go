package usermock

import (
	"fmt"

	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

const RepoErr = "repository error"

type MockRepository struct {
	users  []entity.User
	hasErr bool
}

func NewMockRepository(hasErr bool) *MockRepository {
	users := []entity.User{
		{
			Email:    "test@example.com",
			Password: "$2a$10$pMV1Q1b9jgUQdgWnq4GVOuenS2X.HPns0oMRYmoLdLR1nJL/oONzS", // very_strong_password
		},
	}

	return &MockRepository{
		users:  users,
		hasErr: hasErr,
	}
}

func (m *MockRepository) Register(u entity.User) (*entity.User, error) {
	if m.hasErr {
		return nil, fmt.Errorf(RepoErr)
	}

	u.ID = "new_id"
	m.users = append(m.users, u)

	return &u, nil
}

func (m *MockRepository) GetUserByEmail(email string) (*entity.User, error) {
	if m.hasErr {
		return nil, fmt.Errorf(RepoErr)
	}

	for _, user := range m.users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, richerror.New("MockRepo.GetUserByEmail").WhitMessage(errmsg.ErrAuthUserNotFound)
}

func (m *MockRepository) DoesUserExistsByEmail(email string) (bool, error) {
	if m.hasErr {
		return false, fmt.Errorf(RepoErr)
	}

	for _, user := range m.users {
		if user.Email == email {
			return true, nil
		}
	}

	return false, nil
}
