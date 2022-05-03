package lifestage

import "gocycles/model"

type LifeStage interface {
	Steps() []model.Stage
	Async() bool
}

type lifeStage struct {
	steps []model.Stage
	async bool
}

func NewLifeStage() LifeStage {
	return &lifeStage{}
}

func (l *lifeStage) Steps() []model.Stage {
	return l.steps
}

func (l *lifeStage) Async() bool {
	return l.async
}
