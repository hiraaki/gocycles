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
	for {
		select {
		case state := <-a.State:
			go a.Execute(ctx, state)
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

func (a *App) Execute(ctx context.Context, state int) {
	switch state {
	case 0:
		log.Printf("Start Fase(%v)", state)
		a.runCycle(ctx, a.Lifelifecycle.Start(), state)
	case 1:
		log.Printf("Runner Fase(%v)", state)
		a.runCycle(ctx, a.Lifelifecycle.Run(), state)
	case 2:
		log.Printf("waiting Fase(%v)", state)
		a.runCycle(ctx, a.Lifelifecycle.Wait(), state)
	case 3:
		log.Printf("Restarting Fase(%v)", state)
		a.runCycle(ctx, a.Lifelifecycle.Reset(), state)

	case 4:
		log.Println("Shutingdown")
		os.Exit(0)
	}

}

func (a *App) runCycle(ctx context.Context, lifeStage model.LifeStage, state int) {
	g, ctx := errgroup.WithContext(ctx)
	for _, f := range lifeStage.Stages() {
		f := f
		if !f.Async {
			err := f.Step(ctx)
			if err != nil && f.ResetOnError {
				a.resetApp(err)
				return
			}
			a.Err <- err
			continue
		}
		if f.ResetOnError {
			g.Go(func() error {
				return f.Step(ctx)
			})
		}
		go func() {
			err := f.Step(ctx)
			if err != nil {
				a.Err <- err
			}
		}()
	}
	if err := g.Wait(); err != nil {
		a.resetApp(err)
		go a.Log(err)
		return
	}
	a.next(state)
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

func (a *App) Log(err error) {
	a.Err <- err
}

func (a *App) resetApp(err error) {
	log.Printf("application reseting by: %s", err.Error())
	time.Sleep(time.Second * 3)
	a.State <- 3
}
