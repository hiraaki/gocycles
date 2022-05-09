package main

import (
	"context"
	"gocycles/app"
	"gocycles/exemple/module"
	"gocycles/model"
)

func main() {
	life := model.NewLifeClicle(
		model.WithModule(module.NewServiceModule()),
	)
	app := app.App{
		State:         make(chan int),
		Err:           make(chan error),
		Lifelifecycle: life,
	}
	go app.Start()
	app.Run(context.Background())

}
