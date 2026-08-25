package neutron

type NeutronDiagnostics struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewNeutronDiagnostics(label string) *NeutronDiagnostics {
	return &NeutronDiagnostics{label: label, active: true}
}

func (v *NeutronDiagnostics) Activate() {
	v.active = true
	v.sequence++
}

func (v *NeutronDiagnostics) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *NeutronDiagnostics) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *NeutronDiagnostics) Value() float64 {
	return v.value
}

func (v *NeutronDiagnostics) Active() bool {
	return v.active
}

func (v *NeutronDiagnostics) Sequence() uint64 {
	return v.sequence
}

func (v *NeutronDiagnostics) Label() string {
	return v.label
}

func (v *NeutronDiagnostics) Role() string {
	return "neutron diagnostics"
}
