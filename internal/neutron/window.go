package neutron

// ShutdownRecorder captures the neutron decay trace (the "decay curve") and
// tracks whether that curve has been reliably persisted to durable storage.
// Until the trace is marked durable, the recorded curve is considered an
// incomplete record and must not be used to declare shutdown or to recover a
// safe state.
type ShutdownRecorder struct {
	state      string
	decayTrace []float64
	durable    bool
}

func NewShutdownRecorder() *ShutdownRecorder {
	return &ShutdownRecorder{state: "operating"}
}

// PublishState advances the reactor state. Publishing "shutdown" is only
// meaningful once the decay trace has been reliably recorded (see Durable);
// callers must gate any publish on CanDeclareShutdown.
func (r *ShutdownRecorder) PublishState(state string) {
	r.state = state
}

// RecordDecay appends a neutron decay sample. Each fresh sample invalidates
// the previously persisted curve: until the new trace is re-flushed the
// record is incomplete and the recorder reports not durable.
func (r *ShutdownRecorder) RecordDecay(value float64) {
	r.decayTrace = append(r.decayTrace, value)
	r.durable = false
}

// Durable marks the current decay trace as reliably persisted. It is a
// no-op when there is nothing on the curve to persist, so a durable proof
// always carries at least one decay sample.
func (r *ShutdownRecorder) Durable() {
	if len(r.decayTrace) == 0 {
		return
	}
	r.durable = true
}

// Durable reports whether the neutron decay proof has been reliably written
// to durable storage. Shutdown declaration and safe-state recovery must not
// proceed while this is false.
func (r *ShutdownRecorder) DurableProof() bool {
	return r.durable && len(r.decayTrace) > 0
}

// CanDeclareShutdown reports whether the reactor may be declared shut down:
// the state must be "shutdown" and the neutron decay proof must be present
// and reliably persisted. A non-durable (incomplete) decay record never
// authorizes shutdown.
func (r *ShutdownRecorder) CanDeclareShutdown() bool {
	return r.state == "shutdown" && r.DurableProof()
}

func (r *ShutdownRecorder) TraceLength() int {
	return len(r.decayTrace)
}
