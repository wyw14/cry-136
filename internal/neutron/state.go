package neutron

type CalibrationWindow struct {
	revision uint64
	bound    uint64
	value    float64
	valid    bool
}

func NewCalibrationWindow(revision uint64) *CalibrationWindow {
	return &CalibrationWindow{revision: revision, bound: revision, valid: true}
}

func (w *CalibrationWindow) ApplySample(value float64) {
	w.value = value
}

func (w *CalibrationWindow) UpdateRevision(revision uint64) {
	w.revision = revision
	if revision != w.bound {
		w.valid = true
	}
}

func (w *CalibrationWindow) Valid() bool {
	return w.valid && w.value >= 0
}

func ValidForRevision(sample, current uint64) bool {
	return sample == current
}

func IndependentChannels(channels []string) bool {
	seen := map[string]bool{}
	for _, channel := range channels {
		if seen[channel] {
			return false
		}
		seen[channel] = true
	}
	return len(channels) > 0
}
