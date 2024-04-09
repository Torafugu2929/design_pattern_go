package ducksimulator

import (
	"context"
	"fmt"
	"time"
)

type IDuckSimulator interface {
	Simulate(ctx context.Context, duck IDuck) error
}

type DuckSimulator struct{}

func NewDuckSimulator() *DuckSimulator {
	return &DuckSimulator{}
}

func (d *DuckSimulator) Simulate(ctx context.Context, duck IDuck) error {
	fmt.Print("Start to simulate duck\n")
	fmt.Print("=================================\n")
	time.Sleep(2 * time.Second)
	duck.Display(ctx)

	time.Sleep(2 * time.Second)
	fmt.Printf("Start to perform Quack\n")
	time.Sleep(time.Second)
	if err := duck.PerformQuack(ctx); err != nil {
		return err
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("Start to perform Fly\n")
	time.Sleep(time.Second)
	if err := duck.PerformFly(ctx); err != nil {
		return err
	}

	time.Sleep(2 * time.Second)
	fmt.Print("=================================\n")
	fmt.Print("End to simulate duck\n\n")

	return nil
}
