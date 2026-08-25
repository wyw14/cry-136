package cycle

import (
	"github.com/google/uuid"
	"github.com/wyw14/cry-136/internal/model"
)

type State struct {
	CycleID      uuid.UUID
	Phase        model.Lifecycle
	Epoch        uint64
	TripLatched  bool
	CommandEpoch uint64
}

func NewState() State {
	return State{CycleID: uuid.New(), Phase: model.Shutdown, Epoch: 1, CommandEpoch: 1}
}

func (s *State) Begin() {
	s.Phase = model.Preparing
	s.Epoch++
	s.CommandEpoch++
}

func (s *State) Advance(phase model.Lifecycle) {
	s.Phase = phase
	s.CommandEpoch++
}

func (s *State) CancelStartup() {
	s.Phase = model.Shutdown
	s.CommandEpoch++
}

func (s *State) Trip() {
	s.TripLatched = true
	s.Phase = model.Scrammed
	s.CommandEpoch++
}

func (s *State) ResetForCycle(id uuid.UUID) bool {
	if id != s.CycleID {
		return false
	}
	s.TripLatched = false
	s.Phase = model.Shutdown
	s.CommandEpoch++
	return true
}

func (s State) Snapshot() model.CycleSnapshot {
	return model.CycleSnapshot{ID: s.CycleID, Phase: s.Phase, Epoch: s.Epoch, TripLatched: s.TripLatched, CommandEpoch: s.CommandEpoch}
}

func (s State) CurrentEpoch() uint64 {
	return s.Epoch + 1
}

type TripState struct {
	cycle uint64
	trip uint64
}

func NewTripState() *TripState { return &TripState{} }
func (s *TripState) BeginCycle() uint64 { s.cycle++; return s.cycle }
func (s *TripState) Latch(cycle uint64) { s.trip = cycle }
func (s *TripState) Reset(cycle uint64) { if cycle == s.cycle { s.trip = 0 } }
func (s *TripState) TripLatched() bool { return s.trip != 0 }
