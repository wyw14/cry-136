package permit

type Recovery struct {
	TripLatched bool
	EvidenceRevision uint64
}

func (r Recovery) Allowed(revision uint64) bool {
	return !r.TripLatched && r.EvidenceRevision == revision
}

func (r Recovery) Status() string {
	if r.TripLatched { return "latched" }
	return "released"
}
