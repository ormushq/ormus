package event

import (
	"context"
	"fmt"
	"github.com/ormushq/ormus/event"
)

func (d DB) InsertEvent(ctx context.Context, e event.CoreEvent) (event.CoreEvent, error) {
	//d.adapter.Query()
	return event.CoreEvent{}, fmt.Errorf("dsds")
}
