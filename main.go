package main

import (
	"context"
	"errors"
	"fmt"
	"gocycles/app"
	"gocycles/lifecycle"
	"gocycles/model"
	"gocycles/module"
	"time"
)

func runWithErr(ctx context.Context) error {
	fmt.Println("runnig")
	time.Sleep(time.Second * 1)
	return errors.New("runner Faild")
}

func run(ctx context.Context) error {
	fmt.Println("runnig")
	time.Sleep(time.Second * 1)
	return nil
}

func main() {

	mod := module.NewModule(
		module.WithStart(model.Stage{
			Async: false,
			Step:  run,
		}),
		module.WithRun(model.Stage{
			Async: false,
			Step:  runWithErr,
		}),
		module.WithWait(model.Stage{
			Async: false,
			Step:  run,
		}),
		module.WithReset(model.Stage{
			Async: false,
			Step:  run,
		}),
	)

	mod2 := module.NewModule(
		module.WithStart(model.Stage{
			Async: false,
			Step:  run,
		}),
		module.WithRun(model.Stage{
			Async: false,
			Step:  run,
		}),
		module.WithWait(model.Stage{
			Async: false,
			Step:  run,
		}),
		module.WithReset(model.Stage{
			Async: false,
			Step:  run,
		}),
	)

	life := lifecycle.NewLifeClicle(
		lifecycle.WithModule(mod),
		lifecycle.WithModule(mod2),
	)
	app := app.App{
		State:         make(chan int),
		Err:           make(chan error),
		Async:         true,
		Lifelifecycle: life,
	}
	go app.Start()
	app.Run(context.Background())

}
