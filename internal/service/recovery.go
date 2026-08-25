package service

import "github.com/wyw14/cry-136/internal/journal"

func (r *Runtime) RecoverEvents() map[string]string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return journal.Replay(r.journal.Events())
}

func (r *Runtime) RecoveryReady() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.journal.Count() > 0
}
