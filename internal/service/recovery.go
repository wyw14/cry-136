package service

import "github.com/wyw14/cry-136/internal/journal"

func (r *Runtime) RecoverEvents() map[string]string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return journal.Replay(r.journal.Events())
}

// RecoveryReady reports whether the reactor may be restored to a safe state.
// It requires both a persisted journal and a neutron decay proof that has
// been reliably written (durable): the decay curve must not be an incomplete
// record before experiment unload or safety recovery is permitted.
func (r *Runtime) RecoveryReady() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.journal.Count() > 0 && r.recorder.DurableProof()
}

// RecoverSafeState releases the scram latch and permits experiment unload
// only once the neutron decay proof is reliably persisted. Calling it before
// the decay curve is durable is a no-op and returns false.
func (r *Runtime) RecoverSafeState() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if !r.recorder.DurableProof() {
		return false
	}
	r.scram.Release()
	return true
}
