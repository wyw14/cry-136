package scram

type IntegrationState struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewIntegrationState(label string) *IntegrationState {
	return &IntegrationState{label: label, active: true}
}

func (v *IntegrationState) Activate() {
	v.active = true
	v.sequence++
}

func (v *IntegrationState) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *IntegrationState) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *IntegrationState) Value() float64 {
	return v.value
}

func (v *IntegrationState) Active() bool {
	return v.active
}

func (v *IntegrationState) Sequence() uint64 {
	return v.sequence
}

func (v *IntegrationState) Label() string {
	return v.label
}

func (v *IntegrationState) Role() string {
	return "scram integration"
}
