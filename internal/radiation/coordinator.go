package radiation

type CoordinatorState struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewCoordinatorState(label string) *CoordinatorState {
	return &CoordinatorState{label: label, active: true}
}

func (v *CoordinatorState) Activate() {
	v.active = true
	v.sequence++
}

func (v *CoordinatorState) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *CoordinatorState) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *CoordinatorState) Value() float64 {
	return v.value
}

func (v *CoordinatorState) Active() bool {
	return v.active
}

func (v *CoordinatorState) Sequence() uint64 {
	return v.sequence
}

func (v *CoordinatorState) Label() string {
	return v.label
}

func (v *CoordinatorState) Role() string {
	return "radiation coordination"
}
