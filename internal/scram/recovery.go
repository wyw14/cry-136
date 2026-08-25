package scram

type Recovery struct {
	CycleEpoch uint64
	TripEpoch  uint64
}

func (r Recovery) Current() bool { return r.CycleEpoch == r.TripEpoch }
func (r Recovery) Retired() bool { return r.TripEpoch < r.CycleEpoch }
func (r Recovery) Status() string {
	if r.Current() { return "current" }
	return "retired"
}
