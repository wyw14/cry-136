package rod

type RodTelemetry struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewRodTelemetry(label string) *RodTelemetry {
	return &RodTelemetry{label: label, active: true}
}

func (v *RodTelemetry) Activate() {
	v.active = true
	v.sequence++
}

func (v *RodTelemetry) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *RodTelemetry) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *RodTelemetry) Value() float64 {
	return v.value
}

func (v *RodTelemetry) Active() bool {
	return v.active
}

func (v *RodTelemetry) Sequence() uint64 {
	return v.sequence
}

func (v *RodTelemetry) Label() string {
	return v.label
}

func (v *RodTelemetry) Role() string {
	return "rod telemetry"
}
