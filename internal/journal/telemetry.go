package journal

type JournalTelemetry struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewJournalTelemetry(label string) *JournalTelemetry {
	return &JournalTelemetry{label: label, active: true}
}

func (v *JournalTelemetry) Activate() {
	v.active = true
	v.sequence++
}

func (v *JournalTelemetry) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *JournalTelemetry) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *JournalTelemetry) Value() float64 {
	return v.value
}

func (v *JournalTelemetry) Active() bool {
	return v.active
}

func (v *JournalTelemetry) Sequence() uint64 {
	return v.sequence
}

func (v *JournalTelemetry) Label() string {
	return v.label
}

func (v *JournalTelemetry) Role() string {
	return "journal telemetry"
}
