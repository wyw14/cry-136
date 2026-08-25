package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/wyw14/cry-136/internal/model"
)

func (r *Runtime) CreateOperation(name string) model.Operation {
	r.mu.Lock()
	defer r.mu.Unlock()
	operation := model.Operation{ID: uuid.New(), Name: name, Status: "accepted", CycleID: r.cycle.Snapshot().ID, CreatedAt: time.Now().UTC()}
	r.operations = append(r.operations, operation)
	return operation
}

func (r *Runtime) Operations() []model.Operation {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return append([]model.Operation(nil), r.operations...)
}
