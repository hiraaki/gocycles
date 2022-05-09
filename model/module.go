package model

type Module interface {
	Start() Stage
	Run() Stage
	Wait() Stage
	Reset() Stage
}

type module struct {
	start Stage
	run   Stage
	wait  Stage
	reset Stage
}

type moduleOpt func(m *module)

func NewModule(opts ...moduleOpt) Module {
	var mod module
	for _, opt := range opts {
		opt(&mod)
	}
	return &mod
}

func WithStart(s Stage) moduleOpt {
	return func(m *module) {
		m.start = s
	}
}
func WithRun(s Stage) moduleOpt {
	return func(m *module) {
		m.run = s
	}
}
func WithWait(s Stage) moduleOpt {
	return func(m *module) {
		m.wait = s
	}
}
func WithReset(s Stage) moduleOpt {
	return func(m *module) {
		m.reset = s
	}
}

func (m *module) Start() Stage {
	return m.start
}
func (m *module) Run() Stage {
	return m.run
}
func (m *module) Wait() Stage {
	return m.wait
}
func (m *module) Reset() Stage {
	return m.reset
}
