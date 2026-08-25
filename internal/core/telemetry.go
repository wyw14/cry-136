package core

type CoreTelemetry struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewCoreTelemetry(label string) *CoreTelemetry {
	return &CoreTelemetry{label: label, active: true}
}

func (v *CoreTelemetry) Activate() {
	v.active = true
	v.sequence++
}

func (v *CoreTelemetry) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *CoreTelemetry) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *CoreTelemetry) Value() float64 {
	return v.value
}

func (v *CoreTelemetry) Active() bool {
	return v.active
}

func (v *CoreTelemetry) Sequence() uint64 {
	return v.sequence
}

func (v *CoreTelemetry) Label() string {
	return v.label
}

func (v *CoreTelemetry) Role() string {
	return "core telemetry"
}
