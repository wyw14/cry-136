package neutron

type NeutronMetrics struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewNeutronMetrics(label string) *NeutronMetrics {
	return &NeutronMetrics{label: label, active: true}
}

func (v *NeutronMetrics) Activate() {
	v.active = true
	v.sequence++
}

func (v *NeutronMetrics) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *NeutronMetrics) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *NeutronMetrics) Value() float64 {
	return v.value
}

func (v *NeutronMetrics) Active() bool {
	return v.active
}

func (v *NeutronMetrics) Sequence() uint64 {
	return v.sequence
}

func (v *NeutronMetrics) Label() string {
	return v.label
}

func (v *NeutronMetrics) Role() string {
	return "neutron metrics"
}
