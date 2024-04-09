package ducksimulator

import (
	"context"
	"fmt"
	"time"
)

type IFlyBehavior interface {
	Fly(ctx context.Context) error
}

type FlyBehaviorFlyWithWings struct{}

func NewFlyBehaviorFlyWithWings() *FlyBehaviorFlyWithWings {
	return &FlyBehaviorFlyWithWings{}
}

func (f *FlyBehaviorFlyWithWings) Fly(ctx context.Context) error {
	fmt.Printf("Fly with wings\n")
	return nil
}

type FlyBehaviorDoNothing struct{}

func NewFlyBehaviorDoNothing() *FlyBehaviorDoNothing {
	return &FlyBehaviorDoNothing{}
}

func (f *FlyBehaviorDoNothing) Fly(ctx context.Context) error {
	fmt.Print("............\n")
	return nil
}

type FlyBehaviorUseRocketEngine struct{}

func NewFlyBehaviorUseRocketEngine() *FlyBehaviorUseRocketEngine {
	return &FlyBehaviorUseRocketEngine{}
}

func (f *FlyBehaviorUseRocketEngine) Fly(ctx context.Context) error {
	fmt.Print("3\n")
	time.Sleep(time.Second)
	fmt.Print("2\n")
	time.Sleep(time.Second)
	fmt.Print("1\n")
	time.Sleep(time.Second)
	fmt.Print("Boooooooooooom!!!!!!!!!!!!!!!!\n")
	time.Sleep(2 * time.Second)

	return nil
}
