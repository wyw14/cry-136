package permit

type Experiment struct {
	ID       string
	Position string
	Epoch    uint64
	Approved bool
}

func NewExperiment(id, position string, epoch uint64) Experiment {
	return Experiment{ID: id, Position: position, Epoch: epoch}
}

func (e *Experiment) Approve() { e.Approved = true }
func (e Experiment) Ready(epoch uint64) bool { return e.Approved && e.Epoch == epoch && e.Position != "" }
func (e Experiment) Key() string { return e.ID + ":" + e.Position }
