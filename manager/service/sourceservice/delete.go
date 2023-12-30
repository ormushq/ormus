package sourceservice

import (
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s *Service) DeleteSource(id string) error {
	if err := s.repo.DeleteSource(id); err != nil {
		return richerror.New("DeleteSource").WhitWarpError(err)
	}

	return nil
}
