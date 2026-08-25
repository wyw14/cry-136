package permit

type PermitMetrics struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewPermitMetrics(label string) *PermitMetrics {
	return &PermitMetrics{label: label, active: true}
}

func (v *PermitMetrics) Activate() {
	v.active = true
	v.sequence++
}

func (v *PermitMetrics) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *PermitMetrics) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *PermitMetrics) Value() float64 {
	return v.value
}

func (v *PermitMetrics) Active() bool {
	return v.active
}

func (v *PermitMetrics) Sequence() uint64 {
	return v.sequence
}

func (v *PermitMetrics) Label() string {
	return v.label
}

func (v *PermitMetrics) Role() string {
	return "permit metrics"
}
