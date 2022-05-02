package model

import "context"

type Stage func(ctx context.Context) error
