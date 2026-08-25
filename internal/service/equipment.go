package service

import "github.com/wyw14/cry-136/internal/model"

func (r *Runtime) Equipment() []model.Equipment {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.snapshotLocked().Equipment
}

func (r *Runtime) Interlocks() []model.Interlock {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.snapshotLocked().Interlocks
}

func (r *Runtime) Incidents() []model.Incident {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return append([]model.Incident(nil), r.incidents...)
}
