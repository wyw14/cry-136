package permit

type State struct {
	Revision uint64
	RadiationFresh bool
	RadiationConnected bool
}

func (s State) PermitRevision(revision uint64) bool {
	return s.Revision == revision && s.RadiationFresh && s.RadiationConnected
}

func EvidenceAllowsPower(fresh, connected bool, dose float64) bool {
	return fresh && connected && dose < 1
}

func (s *State) SetRadiation(fresh, connected bool) {
	s.RadiationFresh = fresh
	s.RadiationConnected = connected
	s.Revision++
}
