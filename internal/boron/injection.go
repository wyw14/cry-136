package boron

type Valve struct {
	Open bool
	Timeout bool
}

func (v *Valve) OpenValve() { v.Open = true; v.Timeout = false }
func (v *Valve) FailTimeout() { v.Open = false; v.Timeout = true }
func (v Valve) Healthy() bool { return v.Open && !v.Timeout }
