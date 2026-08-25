package rod

import "testing"

func TestRodInsertionRetryRequiresCurrentBottomProof(t *testing.T) {
	if AcceptBottomAck(8, 7) { t.Fatal("retired bottom acknowledgement accepted") }
	if !AcceptBottomAck(8, 8) { t.Fatal("current bottom acknowledgement rejected") }
}
