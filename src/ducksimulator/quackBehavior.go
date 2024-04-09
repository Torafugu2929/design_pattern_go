package ducksimulator

import (
	"context"
	"fmt"
)

type IQuackBehavior interface {
	Quack(ctx context.Context) error
}

type QuackBehaviorQuack struct{}

func NewQuackBehaviorQuack() *QuackBehaviorQuack {
	return &QuackBehaviorQuack{}
}

func (q *QuackBehaviorQuack) Quack(ctx context.Context) error {
	fmt.Print("Quack\n")
	return nil
}

type QuackBehaviorSqueak struct{}

func NewQuackBehaviorSqueak() *QuackBehaviorSqueak {
	return &QuackBehaviorSqueak{}
}

func (q *QuackBehaviorSqueak) Quack(ctx context.Context) error {
	fmt.Print("Squeak\n")
	return nil
}

type QuackBehaviorMute struct{}

func NewQuackBehaviorMute() *QuackBehaviorMute {
	return &QuackBehaviorMute{}
}

func (q *QuackBehaviorMute) Quack(ctx context.Context) error {
	fmt.Print("............\n")
	return nil
}
