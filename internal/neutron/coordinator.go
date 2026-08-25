package neutron

type ShutdownWindow struct {
	threshold float64
	accum     float64
	connected bool
}

func NewShutdownWindow(threshold float64) *ShutdownWindow {
	return &ShutdownWindow{threshold: threshold, connected: true}
}

func (w *ShutdownWindow) Observe(value float64, connected bool) {
	if !connected {
		w.connected = false
		return
	}
	w.connected = true
	if value <= 0.2 {
		w.accum++
	} else {
		w.accum = 0
	}
}

func (w *ShutdownWindow) Ready() bool {
	return w.connected && w.accum >= w.threshold
}

func (w *ShutdownWindow) Accumulated() float64 {
	return w.accum
}

type Coordinator struct {
	window *ShutdownWindow
	trace  []float64
}

func NewCoordinator() *Coordinator {
	return &Coordinator{window: NewShutdownWindow(3)}
}

func (c *Coordinator) Sample(value float64, connected bool) bool {
	c.window.Observe(value, connected)
	c.trace = append(c.trace, value)
	return c.window.Ready()
}

func (c *Coordinator) Trace() []float64 {
	return append([]float64(nil), c.trace...)
}
