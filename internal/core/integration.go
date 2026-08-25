package core

import "github.com/wyw14/cry-136/internal/model"

type PermitInputs struct {
	Core       model.CoreSnapshot
	RodReady   bool
	Experiment bool
}

func (p PermitInputs) Ready() bool {
	return p.Core.CoolantFlow > 0 && p.Core.RadiationOK && p.RodReady && p.Experiment
}

func (p PermitInputs) Description() string {
	if p.Ready() {
		return "ready"
	}
	return "blocked"
}
