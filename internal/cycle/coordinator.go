package cycle

import (
	"sync"

	"github.com/google/uuid"
	"github.com/wyw14/cry-136/internal/model"
)

type Coordinator struct {
	mu    sync.RWMutex
	state State
}

func NewCoordinator() *Coordinator {
	return &Coordinator{state: NewState()}
}

func (c *Coordinator) Start() model.CycleSnapshot {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.state.Begin()
	c.state.Advance(model.Operating)
	return c.state.Snapshot()
}

func (c *Coordinator) CancelStartup() model.CycleSnapshot {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.state.CancelStartup()
	return c.state.Snapshot()
}

func (c *Coordinator) Trip() model.CycleSnapshot {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.state.Trip()
	return c.state.Snapshot()
}

func (c *Coordinator) Reset(id uuid.UUID) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.state.ResetForCycle(id)
}

func (c *Coordinator) Snapshot() model.CycleSnapshot {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.state.Snapshot()
}

func (c *Coordinator) Epoch() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.state.CurrentEpoch()
}
