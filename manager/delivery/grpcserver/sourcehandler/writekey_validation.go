package sourcehandler

import (
	"context"
	source_proto "github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/manager/service/sourceservice"
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

func (w WriteKeyValidationHandler) IsWriteKeyValid(ctx context.Context, req *source_proto.ValidateWriteKeyReq) (*source_proto.ValidateWriteKeyResp, error) {
	resp, err := w.SourceSvc.IsWriteKeyValid(req)
	if err != nil {
		return nil, nil
	}
	return resp, nil
}
