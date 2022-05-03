package main

import (
	"context"
	"errors"
	"fmt"
	"gocycles/app"
	"gocycles/lifecycle"
	"gocycles/module"
	"time"
)

func runWithErr(ctx context.Context) error {
	fmt.Println("runnig")
	return errors.New("runner Faild")
}

func run(ctx context.Context) error {
	fmt.Println("runnig")
	time.Sleep(time.Second * 5)
	return nil
}

func main() {

	mod := module.NewModule(
		module.WithStart(run),
		module.WithRun(runWithErr),
		module.WithWait(run),
		module.WithReset(run),
	)

	mod2 := module.NewModule(
		module.WithStart(run),
		module.WithRun(run),
		module.WithWait(run),
		module.WithReset(run),
	)

	life := lifecycle.NewLifeClicle(
		lifecycle.WithModule(mod),
		lifecycle.WithModule(mod2),
	)
	app := app.App{
		State:         make(chan int),
		Err:           make(chan error),
		Lifelifecycle: life,
	}
	go app.Start()
	app.Run(context.Background())

}
