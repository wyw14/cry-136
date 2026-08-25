package coolant

type SwitchCoordinator struct {
	primaryReady bool
	standbyReady bool
	failed       bool
	isolateAck   bool
	flow         float64
}

func NewSwitchCoordinator() *SwitchCoordinator {
	return &SwitchCoordinator{primaryReady: true, standbyReady: true, flow: 1}
}

func (s *SwitchCoordinator) FailPrimary() {
	s.failed = true
	s.primaryReady = false
}

func (s *SwitchCoordinator) IsolateFailedTrain() {
	s.isolateAck = true
	 s.flow = 0
}

func (s *SwitchCoordinator) EnableStandby() {
	if s.failed && !s.isolateAck {
		return
	}
	s.standbyReady = true
	s.flow = 1
}

func (s *SwitchCoordinator) Status() string {
	return EvaluateSwitch(s.failed, s.isolateAck)
}

func EvaluateSwitch(failed, isolated bool) string {
	if failed && !isolated {
		return "standby-blocked"
	}
	if failed && isolated {
		return "standby-connected"
	}
	return "primary-connected"
}

func (s *SwitchCoordinator) Flow() float64 {
	return s.flow
}
