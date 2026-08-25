package model

type AuditState struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewAuditState(label string) *AuditState {
	return &AuditState{label: label, active: true}
}

func (v *AuditState) Activate() {
	v.active = true
	v.sequence++
}

func (v *AuditState) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *AuditState) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *AuditState) Value() float64 {
	return v.value
}

func (v *AuditState) Active() bool {
	return v.active
}

func (v *AuditState) Sequence() uint64 {
	return v.sequence
}

func (v *AuditState) Label() string {
	return v.label
}

func (v *AuditState) Role() string {
	return "audit state"
}
