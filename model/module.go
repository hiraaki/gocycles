package model

type Module struct {
	Start Stage
	Run   Stage
	Wait  Stage
	Reset Stage
}

type moduleOpt func(m *Module)

func NewModule(opts ...moduleOpt) *Module {
	var mod Module
	for _, opt := range opts {
		opt(&mod)
	}
	return &mod
}

func WithStart(s Stage) moduleOpt {
	return func(m *Module) {
		m.Start = s
	}
}
func WithRun(s Stage) moduleOpt {
	return func(m *Module) {
		m.Run = s
	}
}
func WithWait(s Stage) moduleOpt {
	return func(m *Module) {
		m.Wait = s
	}
}
func WithReset(s Stage) moduleOpt {
	return func(m *Module) {
		m.Reset = s
	}
}
