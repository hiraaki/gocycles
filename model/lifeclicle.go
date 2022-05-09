package model

type lifeCycleOpt func(l *lifecycle)

type Lifecycle interface {
	Start() LifeStage
	Run() LifeStage
	Wait() LifeStage
	Reset() LifeStage
}

type lifecycle struct {
	start LifeStage
	run   LifeStage
	wait  LifeStage
	reset LifeStage
}

var _ Lifecycle = (*lifecycle)(nil)

func NewLifeClicle(opts ...lifeCycleOpt) Lifecycle {
	l := lifecycle{
		start: NewLifeStage(),
		run:   NewLifeStage(),
		wait:  NewLifeStage(),
		reset: NewLifeStage(),
	}
	for _, opt := range opts {
		opt(&l)
	}
	return &l
}
func WithModule(mod Module) lifeCycleOpt {
	return func(l *lifecycle) {
		l.start.AddStage(mod.Start())
		l.run.AddStage(mod.Run())
		l.wait.AddStage(mod.Wait())
		l.reset.AddStage(mod.Reset())
	}
}

func (l *lifecycle) Start() LifeStage {
	return l.start
}
func (l *lifecycle) Run() LifeStage {
	return l.run
}
func (l *lifecycle) Wait() LifeStage {
	return l.wait
}
func (l *lifecycle) Reset() LifeStage {
	return l.reset
}
