package rod

import (
	"sync"

	"github.com/wyw14/cry-136/internal/model"
)

type Controller struct {
	mu    sync.RWMutex
	state State
}

func NewController() *Controller {
	return &Controller{state: NewState()}
}

func (c *Controller) MoveToBottom(epoch uint64) model.Equipment {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.state.MoveBottom(epoch)
	return c.equipment()
}

func (c *Controller) Withdraw(epoch uint64) model.Equipment {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.state.Withdraw(epoch)
	return c.equipment()
}

func (c *Controller) Ready() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.state.Ready()
}

func (c *Controller) equipment() model.Equipment {
	state := "withdrawn"
	if c.state.Secured {
		state = "secured"
	} else if c.state.Position == 0 {
		state = "bottom"
	}
	return model.Equipment{Name: "control-rods", State: state, Revision: c.state.Epoch}
}

func (c *Controller) Equipment() model.Equipment {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.equipment()
}
