package permit

type PermitTelemetry struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewPermitTelemetry(label string) *PermitTelemetry {
	return &PermitTelemetry{label: label, active: true}
}

func (v *PermitTelemetry) Activate() {
	v.active = true
	v.sequence++
}

func (v *PermitTelemetry) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *PermitTelemetry) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *PermitTelemetry) Value() float64 {
	return v.value
}

func (v *PermitTelemetry) Active() bool {
	return v.active
}

func (v *PermitTelemetry) Sequence() uint64 {
	return v.sequence
}

func (v *PermitTelemetry) Label() string {
	return v.label
}

func (v *PermitTelemetry) Role() string {
	return "permit telemetry"
}
