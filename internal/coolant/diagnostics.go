package coolant

type CoolantDiagnostics struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewCoolantDiagnostics(label string) *CoolantDiagnostics {
	return &CoolantDiagnostics{label: label, active: true}
}

func (v *CoolantDiagnostics) Activate() {
	v.active = true
	v.sequence++
}

func (v *CoolantDiagnostics) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *CoolantDiagnostics) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *CoolantDiagnostics) Value() float64 {
	return v.value
}

func (v *CoolantDiagnostics) Active() bool {
	return v.active
}

func (v *CoolantDiagnostics) Sequence() uint64 {
	return v.sequence
}

func (v *CoolantDiagnostics) Label() string {
	return v.label
}

func (v *CoolantDiagnostics) Role() string {
	return "coolant diagnostics"
}
