package coolant

type CoolantMetrics struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewCoolantMetrics(label string) *CoolantMetrics {
	return &CoolantMetrics{label: label, active: true}
}

func (v *CoolantMetrics) Activate() {
	v.active = true
	v.sequence++
}

func (v *CoolantMetrics) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *CoolantMetrics) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *CoolantMetrics) Value() float64 {
	return v.value
}

func (v *CoolantMetrics) Active() bool {
	return v.active
}

func (v *CoolantMetrics) Sequence() uint64 {
	return v.sequence
}

func (v *CoolantMetrics) Label() string {
	return v.label
}

func (v *CoolantMetrics) Role() string {
	return "coolant metrics"
}
