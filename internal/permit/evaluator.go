package permit

type Evaluator struct {
	neutron bool
	coolant bool
	rod     bool
	radiation bool
}

func NewEvaluator() *Evaluator {
	return &Evaluator{}
}

func (e *Evaluator) SetNeutron(value bool) { e.neutron = value }
func (e *Evaluator) SetCoolant(value bool) { e.coolant = value }
func (e *Evaluator) SetRod(value bool) { e.rod = value }
func (e *Evaluator) SetRadiation(value bool) { e.radiation = value }

func (e *Evaluator) Allowed() bool {
	return e.neutron && e.coolant && e.rod && e.radiation
}

func (e *Evaluator) Summary() string {
	if e.Allowed() { return "enabled" }
	return "blocked"
}
