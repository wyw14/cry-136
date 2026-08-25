package service

import "github.com/wyw14/cry-136/internal/model"

type RecoveryOrder struct {
	steps []string
}

func NewRecoveryOrder() *RecoveryOrder { return &RecoveryOrder{} }

func (o *RecoveryOrder) Recover() []string {
	o.steps = []string{"trip-latch", "rod-control", "cycle-fence"}
	return append([]string(nil), o.steps...)
}

func (o *RecoveryOrder) Valid() bool {
	return len(o.steps) == 3 && o.steps[0] == "trip-latch" && o.steps[1] == "rod-control" && o.steps[2] == "cycle-fence"
}

func (r *Runtime) CyclePhase() model.Lifecycle { return r.cycle.Snapshot().Phase }
