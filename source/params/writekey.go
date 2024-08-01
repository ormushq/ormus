package params

import "time"

type WriteKeyMetaData struct {
	WriteKey   string    `json:"write_key"`
	OwnerID    string    `json:"owner_id"`
	SourceID   string    `json:"source_id"`
	CreatedAt  time.Time `json:"created_at"`
	LastUsedAt time.Time `json:"last_used_at"`
	Status     string    `json:"status"`
}
