package rod

type State struct {
	Position   int
	Epoch      uint64
	Withdrawn  bool
	Secured    bool
}

type InsertionProof struct {
	CycleEpoch uint64
	Command    string
	Bottom     bool
}

func NewState() State {
	return State{Position: 100, Epoch: 1, Withdrawn: true}
}

func (s *State) MoveBottom(epoch uint64) {
	s.Position = 0
	s.Epoch = epoch
	s.Withdrawn = false
	s.Secured = true
}

func (s *State) Withdraw(epoch uint64) {
	s.Position = 100
	s.Epoch = epoch
	s.Withdrawn = true
	s.Secured = false
}

func (s State) Ready() bool {
	return s.Position == 100 && s.Withdrawn
}

func (s State) WorkerStopped() bool {
	return !s.Withdrawn || s.Secured
}

func (s State) ResetTrip(epoch uint64) bool {
	return s.Epoch == epoch && s.Secured
}

func AcceptBottomAck(currentEpoch, ackEpoch uint64) bool {
	return currentEpoch == ackEpoch
}

func ProofMatches(proof InsertionProof, epoch uint64) bool {
	return proof.Bottom && proof.CycleEpoch == epoch && proof.Command == "insert-bottom"
}
