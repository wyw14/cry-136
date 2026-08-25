package irradiation

type RadiationMonitor struct {
	connected bool
	fresh     bool
	dose      float64
}

func NewRadiationMonitor() *RadiationMonitor {
	return &RadiationMonitor{connected: true, fresh: true, dose: 0.02}
}

func (m *RadiationMonitor) Observe(dose float64) {
	m.dose = dose
	m.connected = true
	m.fresh = true
}

func (m *RadiationMonitor) LinkDown() {
	m.connected = false
	m.fresh = false
}

func (m *RadiationMonitor) EvidenceFresh() bool {
	return m.connected && m.fresh
}

func (m *RadiationMonitor) PowerPermit() bool {
	return m.EvidenceFresh() && m.dose < 1
}

func (m *RadiationMonitor) Dose() float64 {
	return m.dose
}
