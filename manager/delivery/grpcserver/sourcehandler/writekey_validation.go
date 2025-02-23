package sourcehandler

import (
	"context"

	source_proto "github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/service/sourceservice"
	"github.com/ormushq/ormus/pkg/richerror"
)

type WriteKeyValidationHandler struct {
	source_proto.UnimplementedIsWriteKeyValidServer
	SourceSvc sourceservice.Service
}

func New(soureSvc sourceservice.Service) *WriteKeyValidationHandler {
	return &WriteKeyValidationHandler{
		SourceSvc: soureSvc,
	}
}

func (w WriteKeyValidationHandler) IsWriteKeyValid(_ context.Context, req *source_proto.ValidateWriteKeyReq) (*source_proto.ValidateWriteKeyResp, error) {
	resp, err := w.SourceSvc.IsWriteKeyValid(req)
	if err != nil {
		logger.L().Error(err.Error())
		return nil, richerror.New("delivery.grpc_server").WithWrappedError(err).WhitKind(richerror.KindUnexpected)
	}

	return resp, nil
}
