package neutron

import "testing"

// RecordDecay must invalidate any previously persisted curve: until the new
// trace is re-flushed the recorder reports an incomplete (non-durable) record.
func TestRecordDecayInvalidatesDurableProof(t *testing.T) {
	recorder := NewShutdownRecorder()

	recorder.RecordDecay(0.1)
	recorder.Durable()
	if !recorder.DurableProof() {
		t.Fatalf("expected durable proof after flushing first decay sample")
	}

	// A fresh sample arrives: the curve is now an incomplete record again until
	// it has been reliably written once more.
	recorder.RecordDecay(0.05)
	if recorder.DurableProof() {
		t.Fatalf("expected incomplete record after new decay sample, before flush")
	}
}

// Durable must be a no-op when there is no decay trace to persist, so a durable
// proof always carries at least one sample (no empty-curve shutdown).
func TestDurableNoOpOnEmptyTrace(t *testing.T) {
	recorder := NewShutdownRecorder()
	recorder.Durable()
	if recorder.DurableProof() {
		t.Fatalf("expected non-durable proof when decay trace is empty")
	}
}

// CanDeclareShutdown must require both the shutdown state and a reliably
// persisted decay curve. A non-durable (incomplete) record never authorizes
// shutdown declaration.
func TestCanDeclareShutdownRequiresDurableProof(t *testing.T) {
	recorder := NewShutdownRecorder()

	// State set to shutdown, but no decay recorded: incomplete record.
	recorder.PublishState("shutdown")
	if recorder.CanDeclareShutdown() {
		t.Fatalf("must not declare shutdown before decay is recorded")
	}

	// Decay recorded but not yet reliably persisted: still incomplete.
	recorder.RecordDecay(0)
	if recorder.CanDeclareShutdown() {
		t.Fatalf("must not declare shutdown while decay record is not durable")
	}

	// Decay reliably persisted: shutdown may now be declared.
	recorder.Durable()
	if !recorder.CanDeclareShutdown() {
		t.Fatalf("expected shutdown to be declarable once decay proof is durable")
	}
}
