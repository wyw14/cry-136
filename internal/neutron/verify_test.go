package neutron

import "testing"

func TestShutdownWindowBreaksOnNeutronTelemetryGap(t *testing.T) {
	w := NewShutdownWindow(2)
	w.Observe(0.1, true)
	w.Observe(0.1, true)
	w.Observe(0, false)
	w.Observe(0.1, true)
	if w.Ready() { t.Fatal("telemetry gap retained shutdown dwell") }
}
