package rod

import "github.com/wyw14/cry-136/internal/model"

type Integration struct {
	CycleEpoch uint64
	Permit     bool
	State      State
}

func (i Integration) Ready() bool {
	return i.Permit && i.State.Epoch == i.CycleEpoch && i.State.Ready()
}

func (i Integration) Equipment() model.Equipment {
	status := "blocked"
	if i.Ready() {
		status = "ready"
	}
	return model.Equipment{Name: "rod-integration", State: status, Revision: i.CycleEpoch}
}
