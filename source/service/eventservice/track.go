package eventservice

import (
	"context"
	"github.com/ormushq/ormus/event"
	"github.com/ormushq/ormus/pkg/richerror"
	"github.com/ormushq/ormus/source/params"
	"time"
)

func (s Service) Track(ctx context.Context, req params.TrackEventRequest) (params.TrackEventResponse, error) {
	event, err := s.repo.InsertEvent(ctx, event.CoreEvent{
		MessageID:         req.MessageID,
		Type:              req.Type,
		Name:              req.Name,
		Properties:        nil,
		Integration:       nil,
		Ctx:               nil,
		SendAt:            time.Time{},
		ReceivedAt:        time.Time{},
		OriginalTimeStamp: time.Time{},
		Timestamp:         time.Time{},
		AnonymousID:       req.AnonymousID,
		UserID:            req.UserID,
		GroupID:           req.GroupID,
		PreviousID:        req.PreviousID,
		Event:             req.Event,
		WriteKey:          req.WriteKey,
		MetaData:          req.MetaData,
		Options:           nil,
	})
	if err != nil {
		return params.TrackEventResponse{}, richerror.New("Track").WithWrappedError(err)
	}

	return params.TrackEventResponse{Event: event}, nil
}
