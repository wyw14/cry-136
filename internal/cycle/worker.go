package cycle

type CycleWorker struct {
	name      string
	started   bool
	cancelled bool
	stopped   bool
	epoch     uint64
}

func NewCycleWorker(name string) *CycleWorker {
	return &CycleWorker{name: name}
}

func (v *CycleWorker) Start() {
	v.started = true
	v.cancelled = false
	v.stopped = false
	v.epoch++
}

func (v *CycleWorker) Cancel() {
	v.cancelled = true
	v.stopped = true
}

func (v *CycleWorker) Stopped() bool {
	return v.started && v.cancelled && v.stopped
}

func (v *CycleWorker) Name() string {
	return v.name
}

func (v *CycleWorker) Epoch() uint64 {
	return v.epoch
}
