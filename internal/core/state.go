package core

import (
	"sync"

	"github.com/wyw14/cry-136/internal/model"
)

type Controller struct {
	mu       sync.RWMutex
	revision uint64
	power    float64
	flow     float64
	trace    []float64
	reading  model.RadiationReading
}

func NewController() *Controller {
	return &Controller{revision: 1, flow: 1.0, reading: model.NewRadiationReading(0.02, true)}
}

func (c *Controller) ApplyPower(power float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.power = power
	c.revision++
}

func (c *Controller) ApplyCooling(flow float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.flow = flow
	c.revision++
}

func (c *Controller) AppendNeutron(value float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.trace) == 12 {
		c.trace = c.trace[1:]
	}
	c.trace = append(c.trace, value)
	c.revision++
}

func (c *Controller) SetRadiation(reading model.RadiationReading) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.reading = reading
	c.revision++
}

func (c *Controller) Snapshot() model.CoreSnapshot {
	c.mu.RLock()
	defer c.mu.RUnlock()
	trace := append([]float64(nil), c.trace...)
	return model.CoreSnapshot{Revision: c.revision, Power: c.power, CoolantFlow: c.flow, NeutronTrace: trace, RadiationOK: c.reading.Connected && c.reading.Fresh, ShutdownReady: len(trace) > 0 && c.flow > 0}
}

func (c *Controller) Revision() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.revision
}

func DwellEligible(flow float64, connected bool) bool {
	return connected && flow > 0
}
