package scram

import "github.com/wyw14/cry-136/internal/model"

type Coordinator struct {
	latched bool
	reason  string
}

func NewCoordinator() *Coordinator { return &Coordinator{} }

func (c *Coordinator) Trigger(reason string) model.Interlock {
	c.latched = true
	c.reason = reason
	return model.Interlock{Name: "scram", Engaged: true, Reason: reason}
}

func (c *Coordinator) Latched() bool { return c.latched }
func (c *Coordinator) Reason() string { return c.reason }

func (c *Coordinator) Release() bool {
	if c.reason == "" { return false }
	c.latched = false
	return true
}
