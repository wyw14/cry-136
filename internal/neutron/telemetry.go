package neutron

type NeutronTelemetry struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewNeutronTelemetry(label string) *NeutronTelemetry {
	return &NeutronTelemetry{label: label, active: true}
}

func (v *NeutronTelemetry) Activate() {
	v.active = true
	v.sequence++
}

func (v *NeutronTelemetry) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *NeutronTelemetry) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *NeutronTelemetry) Value() float64 {
	return v.value
}

func (v *NeutronTelemetry) Active() bool {
	return v.active
}

func (v *NeutronTelemetry) Sequence() uint64 {
	return v.sequence
}

func (v *NeutronTelemetry) Label() string {
	return v.label
}

func (v *NeutronTelemetry) Role() string {
	return "neutron telemetry"
}
