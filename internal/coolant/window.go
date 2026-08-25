package coolant

type Window struct {
	minimum float64
	maximum float64
	value   float64
}

func NewWindow(minimum, maximum float64) *Window {
	return &Window{minimum: minimum, maximum: maximum}
}

func (w *Window) Observe(value float64) {
	w.value = value
}

func (w *Window) Stable() bool {
	return w.value >= w.minimum && w.value <= w.maximum
}

func (w *Window) Value() float64 {
	return w.value
}
