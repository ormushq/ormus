package sourcemock

import (
	"fmt"

	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

const RepoErr = "repository error"

type DefaultSourceTest struct {
	WriteKey    string
	Name        string
	Description string
	ProjectID   string
	OwnerID     string
}

func DefaultSource() DefaultSourceTest {
	return DefaultSourceTest{
		WriteKey:    "write_key",
		Name:        "name",
		Description: "description",
		ProjectID:   "project_id",
		OwnerID:     "owner_id",
	}
}

type MockRepo struct {
	sources []*entity.Source
	hasErr  bool
}

func NewMockRepository(hasErr bool) *MockRepo {
	var sources []*entity.Source
	defaultSource := DefaultSource()
	sources = append(sources,
		&entity.Source{
			WriteKey:    defaultSource.WriteKey,
			Name:        defaultSource.Name,
			Description: defaultSource.Description,
			ProjectID:   defaultSource.ProjectID,
			OwnerID:     defaultSource.OwnerID,
		})

	return &MockRepo{
		sources: sources,
		hasErr:  hasErr,
	}
}

func (m *MockRepo) InsertSource(source *entity.Source) error {
	if m.hasErr {
		return richerror.New("MockRepo.InsertSource").WhitWarpError(fmt.Errorf(RepoErr))
	}

	m.sources = append(m.sources, source)

	return nil
}

func (m *MockRepo) UpdateSource(id string, source *entity.Source) error {
	if m.hasErr {
		return richerror.New("MockRepo.UpdateSource").WhitWarpError(fmt.Errorf(RepoErr))
	}

	for i, s := range m.sources {
		if s.WriteKey == id {
			m.sources[i] = source

			return nil
		}
	}

	return richerror.New("MockRepo.UpdateSource").WhitMessage(errmsg.ErrUserNotFound)
}

func (m *MockRepo) DeleteSource(id string) error {
	if m.hasErr {
		return richerror.New("MockRepo.DeleteSource").WhitWarpError(fmt.Errorf(RepoErr))
	}

	for i, s := range m.sources {
		if s.WriteKey == id {
			m.sources[i] = &entity.Source{}
			return nil
		}
	}

	return richerror.New("MockRepo.DeleteSource").WhitMessage(errmsg.ErrUserNotFound)
}

func (m *MockRepo) GetUserSourceById(ownerID, id string) (*entity.Source, error) {
	if m.hasErr {
		return nil, richerror.New("MockRepo.GetUserSourceById").WhitWarpError(fmt.Errorf(RepoErr))
	}

	for _, s := range m.sources {
		if s.WriteKey == id && s.OwnerID == ownerID {
			return s, nil
		}
	}

	return nil, richerror.New("MockRepo.GetUserSourceById").WhitMessage(errmsg.ErrUserNotFound)
}

func (m *MockRepo) IsSourceAlreadyCreatedByName(name string) (bool, error) {
	if m.hasErr {
		return false, richerror.New("MockRepo.IsSourceAlreadyCreatedByName").WhitWarpError(fmt.Errorf(RepoErr))
	}

	for _, s := range m.sources {
		if s.Name == name {
			return true, nil
		}
	}

	return false, nil
}
