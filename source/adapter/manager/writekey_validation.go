package manager

import (
	"context"

	source_proto "github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/source"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Manager struct {
	config source.Config
}

func New(config source.Config) *Manager {
	return &Manager{config: config}
}

func (m *Manager) IsWriteKeyValid(ctx context.Context, req *source_proto.ValidateWriteKeyReq) (*source_proto.ValidateWriteKeyResp, error) {
	conn, err := grpc.NewClient(m.config.WriteKeyValidationAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := source_proto.NewIsWriteKeyValidClient(conn)
	resp, err := client.IsWriteKeyValid(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
