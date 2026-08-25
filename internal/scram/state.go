package scram

type LatchState struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewLatchState(label string) *LatchState {
	return &LatchState{label: label, active: true}
}

func (v *LatchState) Activate() {
	v.active = true
	v.sequence++
}

func (v *LatchState) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *LatchState) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *LatchState) Value() float64 {
	return v.value
}

func (v *LatchState) Active() bool {
	return v.active
}

func (v *LatchState) Sequence() uint64 {
	return v.sequence
}

func (v *LatchState) Label() string {
	return v.label
}

func (v *LatchState) Role() string {
	return "scram latch"
}
