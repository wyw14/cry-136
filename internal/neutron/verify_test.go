package neutron

import "testing"

func TestCalibrationRevisionInvalidatesPowerWindow(t *testing.T) {
	w := NewCalibrationWindow(4)
	w.ApplySample(0.5)
	w.UpdateRevision(5)
	if w.Valid() { t.Fatal("window accepted sample from retired calibration") }
}
