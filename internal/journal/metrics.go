package journal

type JournalMetrics struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewJournalMetrics(label string) *JournalMetrics {
	return &JournalMetrics{label: label, active: true}
}

func (v *JournalMetrics) Activate() {
	v.active = true
	v.sequence++
}

func (v *JournalMetrics) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *JournalMetrics) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *JournalMetrics) Value() float64 {
	return v.value
}

func (v *JournalMetrics) Active() bool {
	return v.active
}

func (v *JournalMetrics) Sequence() uint64 {
	return v.sequence
}

func (v *JournalMetrics) Label() string {
	return v.label
}

func (v *JournalMetrics) Role() string {
	return "journal metrics"
}
