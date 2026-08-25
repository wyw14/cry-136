package rod

type RodDiagnostics struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewRodDiagnostics(label string) *RodDiagnostics {
	return &RodDiagnostics{label: label, active: true}
}

func (v *RodDiagnostics) Activate() {
	v.active = true
	v.sequence++
}

func (v *RodDiagnostics) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *RodDiagnostics) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *RodDiagnostics) Value() float64 {
	return v.value
}

func (v *RodDiagnostics) Active() bool {
	return v.active
}

func (v *RodDiagnostics) Sequence() uint64 {
	return v.sequence
}

func (v *RodDiagnostics) Label() string {
	return v.label
}

func (v *RodDiagnostics) Role() string {
	return "rod diagnostics"
}
