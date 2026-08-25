package service

type ServiceMetrics struct {
	active bool
	sequence uint64
	label string
	value float64
}

func NewServiceMetrics(label string) *ServiceMetrics {
	return &ServiceMetrics{label: label, active: true}
}

func (v *ServiceMetrics) Activate() {
	v.active = true
	v.sequence++
}

func (v *ServiceMetrics) Deactivate() {
	v.active = false
	v.sequence++
}

func (v *ServiceMetrics) SetValue(value float64) {
	v.value = value
	v.sequence++
}

func (v *ServiceMetrics) Value() float64 {
	return v.value
}

func (v *ServiceMetrics) Active() bool {
	return v.active
}

func (v *ServiceMetrics) Sequence() uint64 {
	return v.sequence
}

func (v *ServiceMetrics) Label() string {
	return v.label
}

func (v *ServiceMetrics) Role() string {
	return "service metrics"
}
