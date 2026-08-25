package core

type RecoveryGuard struct {
	TripLatched bool
	Revision    uint64
}

func (g RecoveryGuard) AllowsControl() bool {
	return !g.TripLatched
}

func (g RecoveryGuard) Matches(revision uint64) bool {
	return g.Revision == revision
}

func RecoverGuard(previous RecoveryGuard, revision uint64) RecoveryGuard {
	if previous.Revision > revision {
		return previous
	}
	previous.Revision = revision
	return previous
}
