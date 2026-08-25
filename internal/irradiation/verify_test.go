package irradiation

import "testing"

func TestRadiationDisconnectCannotBecomeSafeEvidence(t *testing.T) {
	m := NewRadiationMonitor()
	m.LinkDown()
	if m.EvidenceFresh() || m.PowerPermit() { t.Fatal("disconnected monitor became safe evidence") }
}
