package radiation

type RadiationState struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewRadiationState(label string) *RadiationState {
	return &RadiationState{label: label, active: true}
}

func (v *RadiationState) Activate() {
	v.active = true
	v.sequence++
}

func (v *RadiationState) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *RadiationState) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *RadiationState) Value() float64 {
	return v.value
}

func (v *RadiationState) Active() bool {
	return v.active
}

func (v *RadiationState) Sequence() uint64 {
	return v.sequence
}

func (v *RadiationState) Label() string {
	return v.label
}

func (v *RadiationState) Role() string {
	return "radiation state"
}
