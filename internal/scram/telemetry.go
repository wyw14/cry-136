package scram

type ScramTelemetry struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewScramTelemetry(label string) *ScramTelemetry {
	return &ScramTelemetry{label: label, active: true}
}

func (v *ScramTelemetry) Activate() {
	v.active = true
	v.sequence++
}

func (v *ScramTelemetry) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *ScramTelemetry) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *ScramTelemetry) Value() float64 {
	return v.value
}

func (v *ScramTelemetry) Active() bool {
	return v.active
}

func (v *ScramTelemetry) Sequence() uint64 {
	return v.sequence
}

func (v *ScramTelemetry) Label() string {
	return v.label
}

func (v *ScramTelemetry) Role() string {
	return "scram telemetry"
}
