package lifecycle

import (
	"test/model"
	"test/module"
)

type option func(l *lifecycle)

type Lifecycle interface {
	Start() []model.Stage
	Run() []model.Stage
	Wait() []model.Stage
	Reset() []model.Stage
}

type lifecycle struct {
	start []model.Stage
	run   []model.Stage
	wait  []model.Stage
	reset []model.Stage
}

var _ Lifecycle = (*lifecycle)(nil)

func NewLifeClicle(opts ...option) Lifecycle {
	var l lifecycle
	for _, opt := range opts {
		opt(&l)
	}
	return &l
}
func WithModule(mod *module.Module) option {
	return func(l *lifecycle) {
		l.addStart(mod.Start)
		l.addRun(mod.Run)
		l.addWait(mod.Wait)
		l.addReset(mod.Reset)
	}
}

func (l *lifecycle) addStart(s model.Stage) {
	l.start = append(l.start, s)
}
func (l *lifecycle) addRun(s model.Stage) {
	l.run = append(l.run, s)
}
func (l *lifecycle) addWait(s model.Stage) {
	l.wait = append(l.start, s)
}
func (l *lifecycle) addReset(s model.Stage) {
	l.reset = append(l.start, s)
}
func (l *lifecycle) Start() []model.Stage {
	return l.start
}
func (l *lifecycle) Run() []model.Stage {
	return l.run
}
func (l *lifecycle) Wait() []model.Stage {
	return l.wait
}
func (l *lifecycle) Reset() []model.Stage {
	return l.reset
}
