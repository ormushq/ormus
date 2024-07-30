package grpcserver

import (
	"context"
	"errors"
	"github.com/ormushq/ormus/contract/protobuf/manager/goproto/writekey"
	"github.com/ormushq/ormus/manager/repository/sourcerepo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) GetWriteKey(ctx context.Context,
	req *writekey.GetWriteKeyRequest) (*writekey.GetWriteKeyResponse, error) {

	if req.WriteKey == "" {
		return nil, status.Error(codes.InvalidArgument, "write key is required")
	}

	select {
	case <-ctx.Done():
		return nil, status.Error(codes.Canceled, "request was canceled")
	default:
	}

	// Retrieve the write key metadata from your database
	// This is where we call the database access layer
	metadata, err := s.sourceRepo.GetWriteKey(ctx, req.WriteKey)
	if err != nil {
		if errors.Is(err, sourcerepo.ErrWriteKeyNotFound) {
			return nil, status.Error(codes.NotFound, "write key not found")
		}
		return nil, status.Error(codes.Internal, "failed to retrieve write key")
	}

	select {
	case <-ctx.Done():
		return nil, status.Error(codes.Canceled, "request was canceled")
	default:
	}

	protoMetadata := &writekey.WriteKeyMetaData{
		WriteKey:   metadata.WriteKey,
		OwnerId:    metadata.OwnerID,
		SourceId:   metadata.SourceID,
		CreatedAt:  timestamppb.New(metadata.CreatedAt),
		LastUsedAt: timestamppb.New(metadata.LastUsedAt),
		Status:     writekey.WriteKeyStatus(writekey.WriteKeyStatus_value[string(metadata.Status)]),
	}

	return &writekey.GetWriteKeyResponse{Metadata: protoMetadata}, nil
}
