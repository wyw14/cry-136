package boron

type Injector struct {
	failed bool
	message string
}

func NewInjector() *Injector { return &Injector{} }

func (i *Injector) Inject(success bool, message string) {
	i.failed = !success
	i.message = message
}

func (i *Injector) Failed() bool { return i.failed }
func (i *Injector) Error() string { return i.message }

type ScramResult struct {
	Status string
	Error  string
}

func MergeScram(rodSecured bool, injectionError string) ScramResult {
	if injectionError != "" {
		return ScramResult{Status: "failed", Error: injectionError}
	}
	if rodSecured {
		return ScramResult{Status: "secured"}
	}
	return ScramResult{Status: "incomplete", Error: "rod insertion incomplete"}
}
