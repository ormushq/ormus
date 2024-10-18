package sourceservice

import (
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/managerparam/sourceparam"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

func (s Service) Update(req sourceparam.UpdateRequest) (sourceparam.UpdateResponse, error) {
	const op = "sourceService.Update"

	vErr := s.validator.ValidateUpdateRequest(req)
	if vErr != nil {
		return sourceparam.UpdateResponse{}, vErr
	}
	source, err := s.repo.GetWithID(req.SourceID)
	if err != nil {
		return sourceparam.UpdateResponse{}, richerror.New(op).WithWrappedError(err)
	}

	if source.OwnerID != req.UserID {
		return sourceparam.UpdateResponse{}, richerror.New(op).WithKind(richerror.KindForbidden).WithMessage(errmsg.ErrAccessDenied)
	}
	source.Name = req.Name
	source.Description = req.Description
	switch req.Status {
	case string(entity.SourceStatusActive):
		source.Status = entity.SourceStatusActive
	case string(entity.SourceStatusNotActive):
		source.Status = entity.SourceStatusNotActive
	}

	source, err = s.repo.Update(source)
	if err != nil {
		return sourceparam.UpdateResponse{}, richerror.New(op).WithWrappedError(err).WithMessage(errmsg.ErrSomeThingWentWrong)
	}

	return sourceparam.UpdateResponse{
		Source: source,
	}, nil
}
