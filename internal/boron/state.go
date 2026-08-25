package boron

type InjectionState struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewInjectionState(label string) *InjectionState {
	return &InjectionState{label: label, active: true}
}

func (v *InjectionState) Activate() {
	v.active = true
	v.sequence++
}

func (v *InjectionState) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *InjectionState) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *InjectionState) Value() float64 {
	return v.value
}

func (v *InjectionState) Active() bool {
	return v.active
}

func (v *InjectionState) Sequence() uint64 {
	return v.sequence
}

func (v *InjectionState) Label() string {
	return v.label
}

func (v *InjectionState) Role() string {
	return "boron injection"
}
