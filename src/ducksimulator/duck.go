package ducksimulator

import (
	"context"
	"fmt"
)

// Interface

type IDuck interface {
	PerformQuack(ctx context.Context) error
	PerformFly(ctx context.Context) error
	Display(ctx context.Context)
}

// Abstract class
// Duckごとに異なる behavior をインターフェースとして分離する

type AbstractDuck struct {
	flyBehavior   IFlyBehavior
	quackBehavior IQuackBehavior
	name          string
}

func NewAbstractDuck(
	flyBehavior IFlyBehavior,
	quackBehavior IQuackBehavior,
	name string,
) *AbstractDuck {
	return &AbstractDuck{
		flyBehavior:   flyBehavior,
		quackBehavior: quackBehavior,
		name:          name,
	}
}

func (d *AbstractDuck) PerformQuack(ctx context.Context) error {
	err := d.quackBehavior.Quack(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *AbstractDuck) PerformFly(ctx context.Context) error {
	err := d.flyBehavior.Fly(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *AbstractDuck) Display(ctx context.Context) {
	fmt.Printf("Duck name is %s\n", d.name)
}

// 実装。embedding を用い、抽象クラスの継承を再現する
type MalladDuck struct{ *AbstractDuck }

func NewMalladDuck() *MalladDuck {
	return &MalladDuck{
		AbstractDuck: &AbstractDuck{
			flyBehavior:   NewFlyBehaviorFlyWithWings(),
			quackBehavior: NewQuackBehaviorQuack(),
			name:          "MalladDuck",
		},
	}
}

type RubberDuck struct{ *AbstractDuck }

func NewRubberDuck() *RubberDuck {
	return &RubberDuck{
		AbstractDuck: &AbstractDuck{
			flyBehavior:   NewFlyBehaviorDoNothing(),
			quackBehavior: NewQuackBehaviorSqueak(),
			name:          "RubberDuck",
		},
	}
}

type ExtremeDuck struct{ *AbstractDuck }

func NewExtremeDuck() *ExtremeDuck {
	return &ExtremeDuck{
		AbstractDuck: &AbstractDuck{
			flyBehavior:   NewFlyBehaviorUseRocketEngine(),
			quackBehavior: NewQuackBehaviorMute(),
			name:          "ExtremeDuck",
		},
	}
}
