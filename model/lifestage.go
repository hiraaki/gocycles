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
	if opts == nil {
		return &lifeStage{
			stages: []Stage{},
			async:  false,
		}
	}
	return &l
}

func (l *lifeStage) AddStage(s Stage) {
	l.stages = append(l.stages, s)
}

func WithStages(s Stage) lifeStageOpt {
	return func(l *lifeStage) {
		l.stages = append(l.stages, s)
	}
}

func WithAsync(async bool) lifeStageOpt {
	return func(l *lifeStage) {
		l.async = async
	}
}

func (l *lifeStage) Stages() []Stage {
	return l.stages
}

func (l *lifeStage) Async() bool {
	return l.async
}
