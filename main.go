package main

import (
	"context"
	"errors"
	"fmt"
	"gocycles/app"
	"gocycles/model"
	"time"
)

func start(ctx context.Context) error {
	fmt.Println("initializing")
	time.Sleep(time.Second * 1)
	return nil
}

func runWithErr(ctx context.Context) error {
	fmt.Println("runnigErr")
	time.Sleep(time.Second * 1)
	return model.ErrCritical
}

func reseting(ctx context.Context) error {
	fmt.Println("reseting")
	time.Sleep(time.Minute)
	return nil
}

func run(ctx context.Context) error {
	fmt.Println("runnig")
	for {
		select {
		case <-ctx.Done():
			return errors.New("stopping")
		default:
			time.Sleep(time.Second * 3)
			fmt.Println("stilalive")
		}
	}
}

func main() {
	start := model.NewStage(
		model.WithStep(start),
	)
	runWithAsync := model.NewStage(
		model.WithStep(run),
		model.WithAsync(),
	)
	runWithErrAsync := model.NewStage(
		model.WithStep(runWithErr),
		model.WithAsync(),
		model.WithResetOnError(),
	)

	reset := model.NewStage(
		model.WithStep(reseting),
	)

	mod := model.NewModule(
		model.WithStart(start),
		model.WithRun(runWithAsync),
		model.WithReset(reset),
	)

	mod2 := model.NewModule(
		model.WithStart(start),
		model.WithRun(runWithErrAsync),
		model.WithReset(reset),
	)

	life := model.NewLifeClicle(
		model.WithModule(mod),
		model.WithModule(mod2),
	)

	app := app.App{
		State:         make(chan int),
		Err:           make(chan error),
		Lifelifecycle: life,
	}
	go app.Start()
	app.Run(context.Background())

}
