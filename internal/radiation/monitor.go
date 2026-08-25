package radiation

type ChannelState struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewChannelState(label string) *ChannelState {
	return &ChannelState{label: label, active: true}
}

func (v *ChannelState) Activate() {
	v.active = true
	v.sequence++
}

func (v *ChannelState) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *ChannelState) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *ChannelState) Value() float64 {
	return v.value
}

func (v *ChannelState) Active() bool {
	return v.active
}

func (v *ChannelState) Sequence() uint64 {
	return v.sequence
}

func (v *ChannelState) Label() string {
	return v.label
}

func (v *ChannelState) Role() string {
	return "radiation channel"
}
