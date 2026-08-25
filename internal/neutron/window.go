package neutron

type ShutdownRecorder struct {
	state       string
	decayTrace  []float64
	durable     bool
}

func NewShutdownRecorder() *ShutdownRecorder {
	return &ShutdownRecorder{state: "operating"}
}

func (r *ShutdownRecorder) PublishState(state string) {
	r.state = state
}

func (r *ShutdownRecorder) RecordDecay(value float64) {
	r.decayTrace = append(r.decayTrace, value)
	 r.durable = false
}

func (r *ShutdownRecorder) Durable() {
	r.durable = true
}

func (r *ShutdownRecorder) CanDeclareShutdown() bool {
	return r.state == "shutdown" && r.durable && len(r.decayTrace) > 0
}

func (r *ShutdownRecorder) TraceLength() int {
	return len(r.decayTrace)
}
