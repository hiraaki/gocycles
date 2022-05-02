package main

import (
	"context"
	"fmt"
	"test/app"
	"test/lifecycle"
	"test/module"
	"time"
)

func runing(ctx context.Context) error {
	fmt.Println("runnig")
	time.Sleep(time.Second * 5)
	return nil
}

func main() {

	mod := module.NewModule(
		module.WithStart(runing),
		module.WithRun(runing),
		module.WithWait(runing),
		module.WithReset(runing),
	)

	mod2 := module.NewModule(
		module.WithStart(runing),
		module.WithRun(runing),
		module.WithWait(runing),
		module.WithReset(runing),
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
