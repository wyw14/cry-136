package boron

type BoronTelemetry struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewBoronTelemetry(label string) *BoronTelemetry {
	return &BoronTelemetry{label: label, active: true}
}

func (v *BoronTelemetry) Activate() {
	v.active = true
	v.sequence++
}

func (v *BoronTelemetry) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *BoronTelemetry) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *BoronTelemetry) Value() float64 {
	return v.value
}

func (v *BoronTelemetry) Active() bool {
	return v.active
}

func (v *BoronTelemetry) Sequence() uint64 {
	return v.sequence
}

func (v *BoronTelemetry) Label() string {
	return v.label
}

func (v *BoronTelemetry) Role() string {
	return "boron telemetry"
}
