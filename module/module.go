package module

import "test/model"

type Module struct {
	Start model.Stage
	Run   model.Stage
	Wait  model.Stage
	Reset model.Stage
}

type moduleOpt func(m *Module)

func NewModule(opts ...moduleOpt) *Module {
	var mod Module
	for _, opt := range opts {
		opt(&mod)
	}
	return &mod
}

func WithStart(s model.Stage) moduleOpt {
	return func(m *Module) {
		m.Start = s
	}
}
func WithRun(s model.Stage) moduleOpt {
	return func(m *Module) {
		m.Run = s
	}
}
func WithWait(s model.Stage) moduleOpt {
	return func(m *Module) {
		m.Wait = s
	}
}
func WithReset(s model.Stage) moduleOpt {
	return func(m *Module) {
		m.Reset = s
	}
}
