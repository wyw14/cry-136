package service

import "testing"

// TriggerScram must commit the neutron decay proof (sample it and flush it
// durable) before publishing shutdown, so that the decay curve is never an
// incomplete record once the scram snapshot is returned.
func TestTriggerScramCommitsDurableNeutronProof(t *testing.T) {
	runtime := NewRuntime("")

	snapshot := runtime.TriggerScram("operator request")

	if !runtime.recorder.DurableProof() {
		t.Fatalf("neutron decay proof must be durable after scram; decay curve was incomplete")
	}
	if !runtime.recorder.CanDeclareShutdown() {
		t.Fatalf("shutdown must be declarable after scram committed durable neutron proof")
	}
	_ = snapshot
}

// RecoveryReady must require the durable neutron proof: a journal alone (e.g.
// from a cycle-start event) is not sufficient to restore a safe state.
func TestRecoveryReadyRequiresDurableNeutronProof(t *testing.T) {
	runtime := NewRuntime("")

	// StartCycle appends a cycle-start journal event but records no neutron
	// decay proof. Recovery must remain blocked.
	runtime.StartCycle()
	if runtime.RecoveryReady() {
		t.Fatalf("recovery must be blocked before neutron decay proof is durable")
	}

	// After a scram commits the durable proof, recovery becomes ready.
	runtime.TriggerScram("operator request")
	if !runtime.RecoveryReady() {
		t.Fatalf("recovery must be ready after scram commits durable neutron proof")
	}
}

// RecoverSafeState must refuse to release the scram latch (and thereby permit
// experiment unload) until the neutron decay proof is reliably persisted.
func TestRecoverSafeStateBlockedWithoutDurableProof(t *testing.T) {
	runtime := NewRuntime("")

	// No scram yet: no durable proof. Recovery must be refused and must not
	// release the scram latch.
	if runtime.RecoverSafeState() {
		t.Fatalf("must not recover safe state before neutron proof is durable")
	}
	if runtime.scram.Latched() {
		t.Fatalf("scram latch must not be released without durable neutron proof")
	}

	// After scram commits durable proof, safe-state recovery is permitted.
	runtime.TriggerScram("operator request")
	if !runtime.RecoverSafeState() {
		t.Fatalf("must recover safe state once neutron proof is durable")
	}
	if runtime.scram.Latched() {
		t.Fatalf("scram latch must be released after durable safe-state recovery")
	}
}
