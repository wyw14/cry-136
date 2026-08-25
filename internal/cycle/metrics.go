package cycle

type CycleMetrics struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewCycleMetrics(label string) *CycleMetrics {
	return &CycleMetrics{label: label, active: true}
}

func (v *CycleMetrics) Activate() {
	v.active = true
	v.sequence++
}

func (v *CycleMetrics) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *CycleMetrics) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *CycleMetrics) Value() float64 {
	return v.value
}

func (v *CycleMetrics) Active() bool {
	return v.active
}

func (v *CycleMetrics) Sequence() uint64 {
	return v.sequence
}

func (v *CycleMetrics) Label() string {
	return v.label
}

func (v *CycleMetrics) Role() string {
	return "cycle metrics"
}
