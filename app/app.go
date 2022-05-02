package app

import (
	"context"
	"errors"
	"fmt"
	"test/lifecycle"
	"test/model"

	"golang.org/x/sync/errgroup"
)

var ErrCritical error = errors.New("critical failure")

type App struct {
	State         chan int
	Err           chan error
	Lifelifecycle lifecycle.Lifecycle
}

func (a *App) Run(ctx context.Context) {
	for {
		select {
		case state := <-a.State:
			go a.Execute(ctx, state)
		case err := <-a.Err:
			if errors.Is(err, ErrCritical) {
				go a.reset(err)
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
		runCycle(ctx, g, a.Lifelifecycle.Start())
	case 1:
		runCycle(ctx, g, a.Lifelifecycle.Run())
	case 2:
		runCycle(ctx, g, a.Lifelifecycle.Wait())
	case 3:
		runCycle(ctx, g, a.Lifelifecycle.Reset())
	case 4:
		
	}

	a.State <- state + 1

	if err := g.Wait(); err != nil {
		a.Err <- err
	}
}

func runCycle(ctx context.Context, g *errgroup.Group, stages []model.Stage) {
	for _, f := range stages {
		f := f
		g.Go(func() error {
			return f(ctx)
		})
	}
}

func (a *App) reset(err error) {
	fmt.Println(err)
	a.State <- 0
}
