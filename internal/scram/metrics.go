package scram

type ScramMetrics struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewScramMetrics(label string) *ScramMetrics {
	return &ScramMetrics{label: label, active: true}
}

func (v *ScramMetrics) Activate() {
	v.active = true
	v.sequence++
}

func (v *ScramMetrics) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *ScramMetrics) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *ScramMetrics) Value() float64 {
	return v.value
}

func (v *ScramMetrics) Active() bool {
	return v.active
}

func (v *ScramMetrics) Sequence() uint64 {
	return v.sequence
}

func (v *ScramMetrics) Label() string {
	return v.label
}

func (v *ScramMetrics) Role() string {
	return "scram metrics"
}
