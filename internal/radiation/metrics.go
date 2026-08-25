package radiation

type RadiationMetrics struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewRadiationMetrics(label string) *RadiationMetrics {
	return &RadiationMetrics{label: label, active: true}
}

func (v *RadiationMetrics) Activate() {
	v.active = true
	v.sequence++
}

func (v *RadiationMetrics) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *RadiationMetrics) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *RadiationMetrics) Value() float64 {
	return v.value
}

func (v *RadiationMetrics) Active() bool {
	return v.active
}

func (v *RadiationMetrics) Sequence() uint64 {
	return v.sequence
}

func (v *RadiationMetrics) Label() string {
	return v.label
}

func (v *RadiationMetrics) Role() string {
	return "radiation metrics"
}
