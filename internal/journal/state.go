package journal

type State struct {
	TripLatched bool
	Revision    uint64
	CycleEpoch  uint64
}

func NewState() *State {
	return &State{Revision: 1, CycleEpoch: 1}
}

func (s *State) LatchTrip() {
	s.TripLatched = true
	s.Revision++
}

func (s *State) AdvanceCycle() {
	s.CycleEpoch++
	s.Revision++
}

func (s State) RecoveryReady() bool {
	return s.TripLatched && s.Revision > 0
}

func (s State) Current(epoch uint64) bool {
	return s.CycleEpoch == epoch
}
