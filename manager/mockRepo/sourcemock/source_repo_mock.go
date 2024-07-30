package sourcemock

import (
	"fmt"
	writekey "github.com/ormushq/ormus/pkg/write_key"
	"time"

	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/managerparam"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

const RepoErr = "repository error"

type DefaultSourceTest struct {
	ID          string
	WriteKey    entity.WriteKeyMetaData
	Name        string
	Description string
	ProjectID   string
	OwnerID     string
}

func GetDefaultWriteKeyMetaData() entity.WriteKeyMetaData {
	w, _ := writekey.GenerateNewWriteKey()

	return entity.WriteKeyMetaData{
		WriteKey:   w,
		OwnerID:    "owner_id",
		SourceID:   "source_id",
		CreatedAt:  time.Now(),
		LastUsedAt: time.Now(),
		Status:     entity.WriteKeyStatusActive,
	}
}

func DefaultSource() DefaultSourceTest {
	return DefaultSourceTest{
		ID:          "source_id",
		WriteKey:    GetDefaultWriteKeyMetaData(),
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
		ID:               source.ID,
		WriteKeyMetaData: source.WriteKey, // TODO: Only write key value or the whole meta data object?
		Name:             source.Name,
		Description:      source.Description,
		ProjectID:        source.ProjectID,
		OwnerID:          source.OwnerID,
		Status:           source.Status,
		CreateAt:         source.CreateAt,
		UpdateAt:         source.UpdateAt,
		DeleteAt:         source.DeleteAt,
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
				ID:               source.ID,
				WriteKeyMetaData: source.WriteKey,
				Name:             source.Name,
				Description:      source.Description,
				ProjectID:        source.ProjectID,
				OwnerID:          source.OwnerID,
				Status:           source.Status,
				CreateAt:         source.CreateAt,
				UpdateAt:         source.UpdateAt,
				DeleteAt:         source.DeleteAt,
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

func (m *MockRepo) UpdateWriteKeyMetaData(metadata *entity.WriteKeyMetaData) error {
	if m.hasErr {
		return richerror.New("MockRepo.UpdateWriteKey").WithWrappedError(fmt.Errorf(RepoErr))
	}

	for i, s := range m.sources {
		if s.WriteKey.WriteKey == metadata.WriteKey {
			s.WriteKey = *metadata
			m.sources[i] = s
			return nil
		}
	}

	return richerror.New("MockRepo.UpdateWriteKey").WithMessage(errmsg.ErrFailedToUpdateWriteKeyMetaData)
}

func (m *MockRepo) GetWriteKeyMetaData(writeKey string) (*managerparam.WriteKeyMetaData, error) {
	if m.hasErr {
		return nil, richerror.New("MockRepo.GetWriteKey").WithWrappedError(fmt.Errorf(RepoErr))
	}

	for _, s := range m.sources {
		if s.WriteKey.WriteKey == writeKey {
			return &managerparam.WriteKeyMetaData{
				WriteKey:   s.WriteKey.WriteKey,
				OwnerID:    s.WriteKey.OwnerID,
				SourceID:   s.WriteKey.SourceID,
				CreatedAt:  s.WriteKey.CreatedAt,
				LastUsedAt: s.WriteKey.LastUsedAt,
				Status:     managerparam.WriteKeyStatus(s.WriteKey.Status),
			}, nil
		}
	}

	return nil, richerror.New("MockRepo.GetWriteKeyMetaData").WithMessage(errmsg.ErrFailedToGetWriteKeyMetaData)
}

func (m *MockRepo) UpdateLastUsedAt(writeKey string, lastUsedAt time.Time) error {
	if m.hasErr {
		return richerror.New("MockRepo.UpdateLastUsedAt").WithWrappedError(fmt.Errorf(RepoErr))
	}

	for i, s := range m.sources {
		if s.WriteKey.WriteKey == writeKey {
			s.WriteKey.LastUsedAt = lastUsedAt
			m.sources[i] = s
			return nil
		}
	}

	return richerror.New("MockRepo.UpdateLastUsedAt").WithMessage(errmsg.ErrWriteKeyNotFound)
}
