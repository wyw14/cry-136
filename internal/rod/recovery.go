package rod

type Recovery struct {
	Trip bool
	Epoch uint64
}

func (r Recovery) CanControl(epoch uint64) bool {
	return !r.Trip && r.Epoch == epoch
}

func (r Recovery) State() string {
	if r.Trip {
		return "latched"
	}
	return "released"
}
