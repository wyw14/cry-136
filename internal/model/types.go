package model

import (
	"time"

	"github.com/google/uuid"
)

type Lifecycle string

const (
	Shutdown  Lifecycle = "shutdown"
	Preparing Lifecycle = "preparing"
	Critical  Lifecycle = "critical"
	Powering  Lifecycle = "powering"
	Operating Lifecycle = "operating"
	Cooling   Lifecycle = "cooling"
	Scrammed  Lifecycle = "scrammed"
)

type CycleSnapshot struct {
	ID            uuid.UUID `json:"id"`
	Phase         Lifecycle `json:"phase"`
	Epoch         uint64 `json:"epoch"`
	TripLatched   bool      `json:"trip_latched"`
	CommandEpoch  uint64    `json:"command_epoch"`
}

type CoreSnapshot struct {
	Revision      uint64    `json:"revision"`
	Power         float64   `json:"power"`
	CoolantFlow   float64   `json:"coolant_flow"`
	NeutronTrace  []float64 `json:"neutron_trace"`
	RadiationOK   bool      `json:"radiation_ok"`
	ShutdownReady bool      `json:"shutdown_ready"`
}

type Operation struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CycleID   uuid.UUID `json:"cycle_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Equipment struct {
	Name      string `json:"name"`
	State     string `json:"state"`
	Revision  uint64 `json:"revision"`
	Owner     string `json:"owner,omitempty"`
}

type Interlock struct {
	Name      string `json:"name"`
	Engaged   bool   `json:"engaged"`
	Reason    string `json:"reason"`
}

type Incident struct {
	ID        uuid.UUID `json:"id"`
	Severity  string    `json:"severity"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type Snapshot struct {
	Cycle       CycleSnapshot `json:"cycle"`
	Core        CoreSnapshot  `json:"core"`
	Operations  []Operation   `json:"operations"`
	Equipment   []Equipment   `json:"equipment"`
	Interlocks  []Interlock  `json:"interlocks"`
	Incidents   []Incident   `json:"incidents"`
	UpdatedAt   time.Time     `json:"updated_at"`
}
