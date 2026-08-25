package model

import "time"

type RadiationReading struct {
	DoseRate  float64   `json:"dose_rate"`
	Fresh     bool      `json:"fresh"`
	Connected bool      `json:"connected"`
	Observed  time.Time `json:"observed"`
}

type PermitState struct {
	NeutronEvidence bool
	CoolantEvidence bool
	RadiationSafe   bool
	RodReady        bool
	Revision        uint64
}

func (p PermitState) Enabled() bool {
	return p.NeutronEvidence && p.CoolantEvidence && p.RadiationSafe && p.RodReady
}

func (p PermitState) Summary() string {
	if p.Enabled() {
		return "enabled"
	}
	return "blocked"
}

func NewRadiationReading(rate float64, connected bool) RadiationReading {
	return RadiationReading{DoseRate: rate, Fresh: connected, Connected: connected, Observed: time.Now().UTC()}
}
