package irradiation

type IrradiationMetrics struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewIrradiationMetrics(label string) *IrradiationMetrics {
	return &IrradiationMetrics{label: label, active: true}
}

func (v *IrradiationMetrics) Activate() {
	v.active = true
	v.sequence++
}

func (v *IrradiationMetrics) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *IrradiationMetrics) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *IrradiationMetrics) Value() float64 {
	return v.value
}

func (v *IrradiationMetrics) Active() bool {
	return v.active
}

func (v *IrradiationMetrics) Sequence() uint64 {
	return v.sequence
}

func (v *IrradiationMetrics) Label() string {
	return v.label
}

func (v *IrradiationMetrics) Role() string {
	return "irradiation metrics"
}
