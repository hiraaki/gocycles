package model

type LifeStage interface {
	Stages() []Stage
	AddStage(s Stage)
	Async() bool
}

type lifeStage struct {
	stages []Stage
	async  bool
}

type lifeStageOpt func(l *lifeStage)

func NewLifeStage(opts ...lifeStageOpt) LifeStage {
	var l lifeStage
	for _, opt := range opts {
		opt(&l)
	}
	return &l
}

func (l *lifeStage) AddStage(s Stage) {
	if s != nil {
		l.stages = append(l.stages, s)
		if s.Async() {
			l.async = true
		}
	}
}

func WithStages(s Stage) lifeStageOpt {
	return func(l *lifeStage) {
		l.stages = append(l.stages, s)
	}
}

func (l *lifeStage) Stages() []Stage {
	return l.stages
}

func (l *lifeStage) Async() bool {
	return l.async
}
