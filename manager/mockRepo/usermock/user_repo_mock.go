package usermock

import (
	"fmt"

	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

const RepoErr = "repository error"

type DefaultUserTest struct {
	ID       string
	Email    string
	Hash     string
	Password string
}

func DefaultUser() DefaultUserTest {
	return DefaultUserTest{
		ID:       "0000000000",
		Email:    "test@example.com",
		Hash:     "$2a$10$GkuJxcHmx.wsVAVl3V/c3uuj75jWVE5awuxJVfoXzIQBA5zQvl572",
		Password: "HeavYPasS123!",
	}
}

type MockRepository struct {
	users  []entity.User
	hasErr bool
}

func NewMockRepository(hasErr bool) *MockRepository {
	var users []entity.User
	defaultUser := DefaultUser()
	users = append(users, entity.User{ID: defaultUser.ID, Email: defaultUser.Email, Password: defaultUser.Hash})

	return &MockRepository{
		users:  users,
		hasErr: hasErr,
	}
}

func (m *MockRepository) Register(u entity.User) (entity.User, error) {
	if m.hasErr {
		return entity.User{}, fmt.Errorf(RepoErr)
	}

	u.ID = "new_id"
	m.users = append(m.users, u)

	return u, nil
}

func (m *MockRepository) GetUserByEmail(email string) (entity.User, error) {
	if m.hasErr {
		return entity.User{}, fmt.Errorf(RepoErr)
	}

	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}

	return entity.User{}, richerror.New("MockRepo.GetUserByEmail").WithMessage(errmsg.ErrAuthUserNotFound)
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
