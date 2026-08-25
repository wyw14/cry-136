package irradiation

type Lease struct {
	Position string
	Owner    string
	Epoch    uint64
}

func NewLease(position, owner string, epoch uint64) Lease {
	return Lease{Position: position, Owner: owner, Epoch: epoch}
}

func (l Lease) Matches(position, owner string, epoch uint64) bool {
	return l.Position == position && l.Owner == owner && l.Epoch == epoch
}

func (l Lease) Active() bool {
	return l.Position != "" && l.Owner != ""
}
