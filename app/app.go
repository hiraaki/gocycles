package app

import (
	"context"
	"errors"
	"gocycles/lifecycle"
	"gocycles/model"
	"log"
	"os"

	"golang.org/x/sync/errgroup"
)

type App struct {
	State         chan int
	Err           chan error
	Lifelifecycle lifecycle.Lifecycle
	Async         bool
}

func (a *App) Run(ctx context.Context) {
	for {
		select {
		case state := <-a.State:
			if a.Async {
				go a.AsyncExecute(ctx, state)
				continue
			}
			go a.Execute(ctx, state)
		case err := <-a.Err:
			if errors.Is(err, model.ErrCritical) {
				go a.resetState(err)
			}

		}
	}
}
func (a *App) Start() {
	a.State <- 0
}

func (a *App) Execute(ctx context.Context, state int) {
	g, ctx := errgroup.WithContext(ctx)
	switch state {
	case 0:
		log.Printf("Running Start Fase(%v)", state)
		a.runCycle(ctx, g, a.Lifelifecycle.Start())
	case 1:
		log.Printf("Running Runner Fase(%v)", state)
		a.runCycle(ctx, g, a.Lifelifecycle.Run())
	case 2:
		log.Printf("Running waiting Fase(%v)", state)
		a.runCycle(ctx, g, a.Lifelifecycle.Wait())
	case 3:
		log.Printf("Running restarting Fase(%v)", state)
		a.runCycle(ctx, g, a.Lifelifecycle.Reset())

	case 4:
		os.Exit(0)
	}

	if err := g.Wait(); err != nil {
		log.Printf("faze(%v): %s", state, err.Error())
		a.Err <- err
	}
	a.State <- state + 1
}

func (a *App) AsyncExecute(ctx context.Context, state int) {
	g, ctx := errgroup.WithContext(ctx)
	switch state {
	case 0:
		log.Printf("Running Start Fase(%v)", state)
		go a.runCycle(ctx, g, a.Lifelifecycle.Start())
	case 1:
		log.Printf("Running Runner Fase(%v)", state)
		go a.runCycle(ctx, g, a.Lifelifecycle.Run())
	case 2:
		log.Printf("Running waiting Fase(%v)", state)
		go a.runCycle(ctx, g, a.Lifelifecycle.Wait())
	case 3:
		log.Printf("Running restarting Fase(%v)", state)
		go a.runCycle(ctx, g, a.Lifelifecycle.Reset())
	case 4:
		os.Exit(0)
	}

	a.State <- state + 1
	if err := g.Wait(); err != nil {
		log.Printf("faze(%v): %s", state, err.Error())
		a.Err <- err
	}
}

func (a *App) runCycle(ctx context.Context, g *errgroup.Group, stages []model.Stage) {
	for _, f := range stages {
		f := f
		if !f.Async {
			a.Err <- f.Step(ctx)
			continue
		}
		g.Go(func() error {
			return f.Step(ctx)
		})
	}
}

func (a *App) resetState(err error) {
	log.Printf("application reseting by: %s", err.Error())
	a.State <- 0
}
