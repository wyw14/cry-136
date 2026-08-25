package permit

func CoolingPermit(isolated bool, flow float64) bool {
	return isolated && flow >= 1
}

func RodPermit(bottom bool, epoch uint64, current uint64) bool {
	return bottom && epoch == current
}

type Integration struct {
	Cooling bool
	Rod     bool
	Neutron bool
}

func (i Integration) Allowed() bool {
	return i.Cooling && i.Rod && i.Neutron
}
