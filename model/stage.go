package model

import (
	"context"
	"errors"
)

var ErrCritical error = errors.New("critical failure")

type step func(ctx context.Context) error

type stageOpt func(*stage)

type Stage interface {
	Step() step
	Async() bool
	ResetOnError() bool
}

type stage struct {
	step         step
	async        bool
	resetOnError bool
}

var _ Stage = (*stage)(nil)

func NewStage(opts ...stageOpt) Stage {
	var s stage
	for _, opt := range opts {
		opt(&s)
	}
	return &s
}
func WithStep(step step) stageOpt {
	return func(s *stage) {
		s.step = step
	}
}
func WithAsync() stageOpt {
	return func(s *stage) {
		s.async = true
	}
}
func WithResetOnError() stageOpt {
	return func(s *stage) {
		s.resetOnError = true
	}
}
func (s *stage) Step() step {
	return s.step
}
func (s *stage) Async() bool {
	return s.async
}
func (s *stage) ResetOnError() bool {
	return s.resetOnError
}
