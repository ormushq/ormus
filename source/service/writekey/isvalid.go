package writekey

import (
	"context"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/ormushq/ormus/source/params"
)

// IsValid checks whether the writeKey is valid or not.
func (s Service) IsValid(ctx context.Context, writeKey string) (*params.WriteKeyMetaData, error) {
	const op = "writekey.IsValid"

	writeKeyMetaData, err := s.repo.IsValidWriteKey(ctx, writeKey)
	if err != nil {
		return nil, richerror.New(op).WithWrappedError(err)
	}

	// If we got the metadata, then the key is valid
	if writeKeyMetaData != nil {
		dto := &params.WriteKeyMetaData{
			WriteKey:   writeKeyMetaData.WriteKey,
			OwnerID:    writeKeyMetaData.OwnerID,
			SourceID:   writeKeyMetaData.SourceID,
			CreatedAt:  writeKeyMetaData.CreatedAt,
			LastUsedAt: writeKeyMetaData.LastUsedAt,
			Status:     string(writeKeyMetaData.Status),
		}

		return dto, nil
	}

	// no metadata, no error then key is invalid
	return nil, nil
}
