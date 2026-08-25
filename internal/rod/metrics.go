package rod

type RodMetrics struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewRodMetrics(label string) *RodMetrics {
	return &RodMetrics{label: label, active: true}
}

func (v *RodMetrics) Activate() {
	v.active = true
	v.sequence++
}

func (v *RodMetrics) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *RodMetrics) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *RodMetrics) Value() float64 {
	return v.value
}

func (v *RodMetrics) Active() bool {
	return v.active
}

func (v *RodMetrics) Sequence() uint64 {
	return v.sequence
}

func (v *RodMetrics) Label() string {
	return v.label
}

func (v *RodMetrics) Role() string {
	return "rod metrics"
}
