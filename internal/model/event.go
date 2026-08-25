package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID        uuid.UUID `json:"id"`
	Kind      string    `json:"kind"`
	CycleID   uuid.UUID `json:"cycle_id"`
	Payload   string    `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}

func NewEvent(kind string, cycleID uuid.UUID, payload string) Event {
	return Event{ID: uuid.New(), Kind: kind, CycleID: cycleID, Payload: payload, CreatedAt: time.Now().UTC()}
}

func (e Event) String() string {
	return fmt.Sprintf("%s:%s:%s", e.ID, e.Kind, e.Payload)
}
