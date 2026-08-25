package core

type CoreDiagnostics struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewCoreDiagnostics(label string) *CoreDiagnostics {
	return &CoreDiagnostics{label: label, active: true}
}

func (v *CoreDiagnostics) Activate() {
	v.active = true
	v.sequence++
}

func (v *CoreDiagnostics) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *CoreDiagnostics) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *CoreDiagnostics) Value() float64 {
	return v.value
}

func (v *CoreDiagnostics) Active() bool {
	return v.active
}

func (v *CoreDiagnostics) Sequence() uint64 {
	return v.sequence
}

func (v *CoreDiagnostics) Label() string {
	return v.label
}

func (v *CoreDiagnostics) Role() string {
	return "core diagnostics"
}
