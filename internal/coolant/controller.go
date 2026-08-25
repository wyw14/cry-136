package coolant

type Controller struct {
	flow     float64
	reserved float64
	active   string
}

func NewController() *Controller {
	return &Controller{flow: 1, active: "primary"}
}

func (c *Controller) SetFlow(flow float64) {
	c.flow = flow
}

func (c *Controller) Select(train string) {
	c.active = train
}

func (c *Controller) Reserve(flow float64) bool {
	if flow <= 0 || flow > c.flow {
		return false
	}
	c.reserved = flow
	return true
}

func (c *Controller) Snapshot() (string, float64, float64) {
	return c.active, c.flow, c.reserved
}
