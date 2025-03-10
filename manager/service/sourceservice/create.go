package sourceservice

import (
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/managerparam/projectparam"
	"github.com/ormushq/ormus/manager/managerparam/sourceparam"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
	writekey "github.com/ormushq/ormus/pkg/write_key"
)

func (s Service) CreateSource(req sourceparam.CreateRequest) (sourceparam.CreateResponse, error) {
	const op = "sourceService.Create"

	vErr := s.validator.ValidateCreateRequest(req)
	if vErr != nil {
		return sourceparam.CreateResponse{}, vErr
	}

	w, err := writekey.GenerateNewWriteKey()
	if err != nil {
		return sourceparam.CreateResponse{}, err
	}

	getProjectReq := projectparam.GetRequest{
		UserID:    req.UserID,
		ProjectID: req.ProjectID,
	}
	_, err = s.projectSvc.Get(getProjectReq)
	if err != nil {
		return sourceparam.CreateResponse{}, richerror.New(op).WithKind(richerror.KindNotFound).WithMessage(errmsg.ErrProjectNotFound)
	}

	source := entity.Source{
		WriteKey:    entity.WriteKey(w),
		Name:        req.Name,
		Description: req.Description,
		OwnerID:     req.UserID,
		ProjectID:   req.ProjectID,
	}

	source, err = s.repo.Create(source)
	if err != nil {
		logger.L().Error(err.Error())

		return sourceparam.CreateResponse{}, richerror.New(op).WithWrappedError(err).WithMessage(errmsg.ErrSomeThingWentWrong)
	}

	return sourceparam.CreateResponse{Source: source}, nil
}
