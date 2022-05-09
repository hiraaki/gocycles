package app

import (
	"context"
	"gocycles/model"
	"log"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

type App struct {
	State         chan int
	Err           chan error
	Lifelifecycle model.Lifecycle
}

func (a *App) Run(ctx context.Context) {
	g, ctx := errgroup.WithContext(ctx)
	for {
		select {
		case state := <-a.State:
			g.Go(func() error {
				a.Execute(ctx, g, state)
				return nil
			})
		case err := <-a.Err:
			if err != nil {
				log.Println(err)
			}
		}
	}
}
func (a *App) Start() {
	a.State <- 0
}

func (a *App) Execute(ctx context.Context, g *errgroup.Group, state int) {
	switch state {
	case 0:
		log.Printf("Start Fase(%v)", state)
		a.runCycle(ctx, a.Lifelifecycle.Start(), state)
	case 1:
		log.Printf("Runner Fase(%v)", state)
		a.runCycle(ctx, a.Lifelifecycle.Run(), state)
	case 2:
		log.Printf("Waiting to finish(%v)", state)
		a.wait(g)
	case 3:
		log.Printf("Restarting Fase(%v)", state)
		a.runCycle(ctx, a.Lifelifecycle.Reset(), state)
	case 4:
		log.Println("Shutingdown")
		os.Exit(0)
	}
	go a.next(state)
}

func (a *App) runCycle(ctx context.Context, lifeStage model.LifeStage, state int) {
	g, ctx := errgroup.WithContext(ctx)
	for _, f := range lifeStage.Stages() {
		f := f
		if !f.Async() {
			err := f.Step()(ctx)
			a.Err <- err
			if err != nil && f.ResetOnError() {
				go a.resetApp(err)
				return
			}
			continue
		}
		g.Go(func() error {
			err := f.Step()(ctx)
			if f.ResetOnError() {
				return err
			}
			return nil
		})
	}
	if lifeStage.Async() {
		go a.next(state)
	}
	a.wait(g)
}

func (a *App) next(state int) {
	if state < 2 {
		a.State <- state + 1
	}
	if state == 2 {
		a.State <- 4
	}
	if state == 3 {
		a.State <- 0
	}
}

func (a *App) resetApp(err error) {
	log.Printf("application reseting by: %s", err.Error())
	time.Sleep(time.Second * 3)
	a.State <- 3
}

func (a *App) wait(g *errgroup.Group) {
	if err := g.Wait(); err != nil {
		go a.resetApp(err)
		a.Err <- err
		return
	}
}
