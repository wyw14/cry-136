package irradiation

type IrradiationTelemetry struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewIrradiationTelemetry(label string) *IrradiationTelemetry {
	return &IrradiationTelemetry{label: label, active: true}
}

func (v *IrradiationTelemetry) Activate() {
	v.active = true
	v.sequence++
}

func (v *IrradiationTelemetry) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *IrradiationTelemetry) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *IrradiationTelemetry) Value() float64 {
	return v.value
}

func (v *IrradiationTelemetry) Active() bool {
	return v.active
}

func (v *IrradiationTelemetry) Sequence() uint64 {
	return v.sequence
}

func (v *IrradiationTelemetry) Label() string {
	return v.label
}

func (v *IrradiationTelemetry) Role() string {
	return "irradiation telemetry"
}
