package boron

type BoronMetrics struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewBoronMetrics(label string) *BoronMetrics {
	return &BoronMetrics{label: label, active: true}
}

func (v *BoronMetrics) Activate() {
	v.active = true
	v.sequence++
}

func (v *BoronMetrics) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *BoronMetrics) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *BoronMetrics) Value() float64 {
	return v.value
}

func (v *BoronMetrics) Active() bool {
	return v.active
}

func (v *BoronMetrics) Sequence() uint64 {
	return v.sequence
}

func (v *BoronMetrics) Label() string {
	return v.label
}

func (v *BoronMetrics) Role() string {
	return "boron metrics"
}
