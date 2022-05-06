package model

import (
	"context"
	"errors"
)

var ErrCritical error = errors.New("critical failure")

type step func(ctx context.Context) error

type Stage struct {
	Step         step
	Async        bool
	ResetOnError bool
}
