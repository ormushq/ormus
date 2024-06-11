package sourcemock

import (
	"fmt"

	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/managerparam"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

const RepoErr = "repository error"

type DefaultSourceTest struct {
	ID          string
	WriteKey    entity.WriteKey
	Name        string
	Description string
	ProjectID   string
	OwnerID     string
}

func DefaultSource() DefaultSourceTest {
	return DefaultSourceTest{
		ID:          "source_id",
		WriteKey:    entity.WriteKey("writekey"),
		Name:        "name name",
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
			ID:          defaultSource.ID,
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

func (m *MockRepo) InsertSource(source *entity.Source) (*managerparam.AddSourceResponse, error) {
	if m.hasErr {
		return nil, richerror.New("MockRepo.InsertSource").WithWrappedError(fmt.Errorf(RepoErr))
	}

	m.sources = append(m.sources, source)

	return &managerparam.AddSourceResponse{
		ID:          source.ID,
		WriteKey:    string(source.WriteKey),
		Name:        source.Name,
		Description: source.Description,
		ProjectID:   source.ProjectID,
		OwnerID:     source.OwnerID,
		Status:      source.Status,
		CreateAt:    source.CreateAt,
		UpdateAt:    source.UpdateAt,
		DeleteAt:    source.DeleteAt,
	}, nil
}

func (m *MockRepo) UpdateSource(id string, source *entity.Source) (*managerparam.UpdateSourceResponse, error) {
	if m.hasErr {
		return nil, richerror.New("MockRepo.UpdateSource").WithWrappedError(fmt.Errorf(RepoErr))
	}

	for i, s := range m.sources {
		if s.ID == id {
			m.sources[i] = source

			return &managerparam.UpdateSourceResponse{
				ID:          source.ID,
				WriteKey:    string(source.WriteKey),
				Name:        source.Name,
				Description: source.Description,
				ProjectID:   source.ProjectID,
				OwnerID:     source.OwnerID,
				Status:      source.Status,
				CreateAt:    source.CreateAt,
				UpdateAt:    source.UpdateAt,
				DeleteAt:    source.DeleteAt,
			}, nil
		}
	}

	return nil, richerror.New("MockRepo.UpdateSource").WithMessage(errmsg.ErrUserNotFound)
}

func (m *MockRepo) DeleteSource(id, _ string) error {
	if m.hasErr {
		return richerror.New("MockRepo.DeleteSource").WithWrappedError(fmt.Errorf(RepoErr))
	}

	for i, s := range m.sources {
		if s.ID == id {
			m.sources[i] = &entity.Source{}

			return nil
		}
	}

	return richerror.New("MockRepo.DeleteSource").WithMessage(errmsg.ErrUserNotFound)
}

func (m *MockRepo) GetUserSourceByID(ownerID, id string) (*entity.Source, error) {
	if m.hasErr {
		return nil, richerror.New("MockRepo.GetUserSourceById").WithWrappedError(fmt.Errorf(RepoErr))
	}

	for _, s := range m.sources {
		if s.ID == id && s.OwnerID == ownerID {
			return s, nil
		}
	}

	return nil, richerror.New("MockRepo.GetUserSourceById").WithMessage(errmsg.ErrUserNotFound)
}

func (m *MockRepo) IsSourceAlreadyCreatedByName(name string) (bool, error) {
	if m.hasErr {
		return false, richerror.New("MockRepo.IsSourceAlreadyCreatedByName").WithWrappedError(fmt.Errorf(RepoErr))
	}

	for _, s := range m.sources {
		if s.Name == name {
			return true, nil
		}
	}

	return false, nil
}
