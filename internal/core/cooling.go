package core

type CoolingPlan struct {
	PrimaryFlow   float64
	StandbyFlow   float64
	PrimaryReady  bool
	StandbyReady  bool
}

func NewCoolingPlan(primary, standby float64) CoolingPlan {
	return CoolingPlan{PrimaryFlow: primary, StandbyFlow: standby, PrimaryReady: primary > 0, StandbyReady: standby > 0}
}

func (p CoolingPlan) Stable() bool {
	return p.PrimaryReady && p.PrimaryFlow > 0
}

func (p CoolingPlan) StandbyAvailable() bool {
	return p.StandbyReady && p.StandbyFlow > 0
}

func (p CoolingPlan) TotalFlow() float64 {
	return p.PrimaryFlow + p.StandbyFlow
}
