package module

import (
	"gocycles/exemple/service"
	"gocycles/model"
)

type serviceMod struct {
	wsServer service.WSServer
}

func NewServiceModule() model.Module {
	return &serviceMod{
		wsServer: service.NewWSServer(),
	}
}
func (s *serviceMod) Start() model.Stage {
	return model.NewStage(
		model.WithStep(s.wsServer.Start),
		model.WithResetOnError(),
	)
}
func (s *serviceMod) Run() model.Stage {
	return model.NewStage(
		model.WithStep(s.wsServer.Run),
		model.WithResetOnError(),
		model.WithAsync(),
	)
}
func (s *serviceMod) Wait() model.Stage {
	return model.NewStage()
}
func (s *serviceMod) Reset() model.Stage {
	return model.NewStage(
		model.WithStep(s.wsServer.Restart),
		model.WithResetOnError(),
	)
}
