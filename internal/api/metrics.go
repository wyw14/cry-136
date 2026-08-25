package api

type APIMetrics struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewAPIMetrics(label string) *APIMetrics {
	return &APIMetrics{label: label, active: true}
}

func (v *APIMetrics) Activate() {
	v.active = true
	v.sequence++
}

func (v *APIMetrics) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *APIMetrics) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *APIMetrics) Value() float64 {
	return v.value
}

func (v *APIMetrics) Active() bool {
	return v.active
}

func (v *APIMetrics) Sequence() uint64 {
	return v.sequence
}

func (v *APIMetrics) Label() string {
	return v.label
}

func (v *APIMetrics) Role() string {
	return "api metrics"
}
