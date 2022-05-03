package model

import (
	"context"
	"errors"
)

var ErrCritical error = errors.New("critical failure")

type Stage func(ctx context.Context) error
