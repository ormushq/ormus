package managerparam

import "time"

type WriteKeyStatus string

const (
	WriteKeyStatusInactive WriteKeyStatus = "WRITE_KEY_STATUS_INACTIVE"
	WriteKeyStatusActive   WriteKeyStatus = "WRITE_KEY_STATUS_ACTIVE"
)

type WriteKeyMetaData struct {
	WriteKey   string         `json:"write_key"`
	OwnerID    string         `json:"owner_id"`
	SourceID   string         `json:"source_id"`
	CreatedAt  time.Time      `json:"created_at"`
	LastUsedAt time.Time      `json:"last_used_at"`
	Status     WriteKeyStatus `json:"status"`
}
