package core

type CoreMetrics struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewCoreMetrics(label string) *CoreMetrics {
	return &CoreMetrics{label: label, active: true}
}

func (v *CoreMetrics) Activate() {
	v.active = true
	v.sequence++
}

func (v *CoreMetrics) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *CoreMetrics) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *CoreMetrics) Value() float64 {
	return v.value
}

func (v *CoreMetrics) Active() bool {
	return v.active
}

func (v *CoreMetrics) Sequence() uint64 {
	return v.sequence
}

func (v *CoreMetrics) Label() string {
	return v.label
}

func (v *CoreMetrics) Role() string {
	return "core metrics"
}
