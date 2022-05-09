package service

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type WSServer interface {
	Start(ctx context.Context) error
	Run(ctx context.Context) error
	Restart(ctx context.Context) error
}

type wsServer struct {
}

func NewWSServer() WSServer {
	return &wsServer{}
}

func (ws *wsServer) Start(ctx context.Context) error {
	fmt.Println("initializing")
	time.Sleep(time.Second * 1)
	return nil
}

func (ws *wsServer) Run(ctx context.Context) error {
	fmt.Println("runnig")
	for {
		select {
		case <-ctx.Done():
			return errors.New("stopping")
		default:
			time.Sleep(time.Second * 3)
			fmt.Println("stilalive")
		}
	}
}

func (ws *wsServer) Restart(ctx context.Context) error {
	fmt.Println("reseting")
	time.Sleep(time.Minute)
	return nil
}
