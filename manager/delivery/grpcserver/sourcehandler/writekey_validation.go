package sourcehandler

import (
	"context"
	source_proto "github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/manager/service/sourceservice"
)

type WriteKeyValidation struct {
	source_proto.UnimplementedIsWriteKeyValidServer
	SourceSvc sourceservice.Service
}

func (w WriteKeyValidation) IsWriteKeyValid(ctx context.Context, req *source_proto.ValidateWriteKeyReq) (*source_proto.ValidateWriteKeyResp, error) {
	resp, err := w.SourceSvc.IsWriteKeyValid(req)
	if err != nil {
		return nil, nil
	}
	return resp, nil
}
