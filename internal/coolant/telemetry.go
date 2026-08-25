package coolant

type CoolantTelemetry struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewCoolantTelemetry(label string) *CoolantTelemetry {
	return &CoolantTelemetry{label: label, active: true}
}

func (v *CoolantTelemetry) Activate() {
	v.active = true
	v.sequence++
}

func (v *CoolantTelemetry) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *CoolantTelemetry) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *CoolantTelemetry) Value() float64 {
	return v.value
}

func (v *CoolantTelemetry) Active() bool {
	return v.active
}

func (v *CoolantTelemetry) Sequence() uint64 {
	return v.sequence
}

func (v *CoolantTelemetry) Label() string {
	return v.label
}

func (v *CoolantTelemetry) Role() string {
	return "coolant telemetry"
}
