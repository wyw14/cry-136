package cycle

import "github.com/wyw14/cry-136/internal/model"

func CanAdvance(from, to model.Lifecycle) bool {
	if from == model.Scrammed {
		return to == model.Shutdown
	}
	if from == model.Shutdown {
		return to == model.Preparing
	}
	if from == model.Preparing {
		return to == model.Critical || to == model.Shutdown
	}
	if from == model.Critical {
		return to == model.Powering || to == model.Scrammed
	}
	if from == model.Powering {
		return to == model.Operating || to == model.Scrammed
	}
	if from == model.Operating {
		return to == model.Cooling || to == model.Scrammed
	}
	if from == model.Cooling {
		return to == model.Shutdown
	}
	return false
}

func Stable(phase model.Lifecycle) bool {
	return phase == model.Operating || phase == model.Cooling || phase == model.Shutdown
}
