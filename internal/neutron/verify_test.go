package neutron

import "testing"

func TestShutdownWaitsForDurableNeutronDecayTrace(t *testing.T) {
	r := NewShutdownRecorder()
	r.PublishState("shutdown")
	r.RecordDecay(0.2)
	if r.CanDeclareShutdown() { t.Fatal("shutdown declared before durable trace") }
	r.Durable()
	if !r.CanDeclareShutdown() { t.Fatal("durable trace not accepted") }
}
