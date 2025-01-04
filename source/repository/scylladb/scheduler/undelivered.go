package scheduler

import (
	"time"

	proto "github.com/ormushq/ormus/contract/go/source"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/source/repository/scylladb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func init() {
	statements["get_undelivered"] = scylladb.Statement{
		Query:  `SELECT id,created_at,event,name,properties,received_at,send_at,timestamp,type,updated_at,write_key FROM event WHERE delivered = false allow filtering;`,
		Values: []string{},
	}
}

func (r Repository) GetAllUnDeliveredEvents() ([]*proto.NewEvent, error) {
	query, err := r.db.GetStatement(statements["get_undelivered"])
	if err != nil {
		logger.L().Error(err.Error())

		return nil, err
	}
	var rawResults []struct {
		ID         string            `db:"id"`
		CreatedAt  time.Time         `db:"created_at"`
		Event      string            `db:"event"`
		Name       string            `db:"name"`
		Properties map[string]string `db:"properties"`
		ReceivedAt time.Time         `db:"received_at"`
		SendAt     time.Time         `db:"send_at"`
		Timestamp  time.Time         `db:"timestamp"`
		Type       string            `db:"type"`
		UpdatedAt  time.Time         `db:"updated_at"`
		WriteKey   string            `db:"write_key"`
	}
	err = query.Select(&rawResults)
	if err != nil {
		logger.L().Error(err.Error())

		return nil, err
	}
	result := make([]*proto.NewEvent, 0, len(rawResults))
	for _, row := range rawResults {
		event := &proto.NewEvent{
			Id:         row.ID,
			CreatedAt:  timestamppb.New(row.CreatedAt),
			Event:      row.Event,
			Name:       row.Name,
			Properties: row.Properties,
			ReceivedAt: timestamppb.New(row.ReceivedAt),
			SendAt:     timestamppb.New(row.SendAt),
			Timestamp:  timestamppb.New(row.Timestamp),
			Type:       row.Type,
			UpdatedAt:  timestamppb.New(row.UpdatedAt),
			WriteKey:   row.WriteKey,
		}
		result = append(result, event)
	}

	return result, nil
}
