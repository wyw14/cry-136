package rod

type Coordinator struct {
	controller *Controller
	lastEpoch  uint64
}

func NewCoordinator(controller *Controller) *Coordinator {
	return &Coordinator{controller: controller}
}

func (c *Coordinator) Insert(epoch uint64) bool {
	if epoch < c.lastEpoch {
		return false
	}
	c.lastEpoch = epoch
	c.controller.MoveToBottom(epoch)
	return true
}

func (c *Coordinator) Retract(epoch uint64) bool {
	if epoch < c.lastEpoch {
		return false
	}
	c.lastEpoch = epoch
	c.controller.Withdraw(epoch)
	return true
}

func (c *Coordinator) Epoch() uint64 {
	return c.lastEpoch
}
