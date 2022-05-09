package main

import (
	"context"
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
	time.Sleep(time.Second * 10)
	return nil
}

func run(ctx context.Context) error {
	fmt.Println("runnig")
	time.Sleep(time.Second * 1)
	return nil
}

func main() {

	mod := model.NewModule(
		model.WithStart(model.Stage{
			Async: false,
			Step:  start,
		}),
		model.WithRun(model.Stage{
			Async:        true,
			Step:         runWithErr,
			ResetOnError: true,
		}),
		model.WithWait(model.Stage{
			Async: false,
			Step:  run,
		}),
		model.WithReset(model.Stage{
			Async: false,
			Step:  reseting,
		}),
	)

	mod2 := model.NewModule(
		model.WithStart(model.Stage{
			Async: false,
			Step:  start,
		}),
		model.WithRun(model.Stage{
			Async: true,
			Step:  run,
		}),
		model.WithWait(model.Stage{
			Async: false,
			Step:  run,
		}),
		model.WithReset(model.Stage{
			Async: false,
			Step:  reseting,
		}),
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
