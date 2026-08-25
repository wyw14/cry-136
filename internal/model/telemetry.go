package model

type TelemetryState struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewTelemetryState(label string) *TelemetryState {
	return &TelemetryState{label: label, active: true}
}

func (v *TelemetryState) Activate() {
	v.active = true
	v.sequence++
}

func (v *TelemetryState) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *TelemetryState) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *TelemetryState) Value() float64 {
	return v.value
}

func (v *TelemetryState) Active() bool {
	return v.active
}

func (v *TelemetryState) Sequence() uint64 {
	return v.sequence
}

func (v *TelemetryState) Label() string {
	return v.label
}

func (v *TelemetryState) Role() string {
	return "telemetry"
}
