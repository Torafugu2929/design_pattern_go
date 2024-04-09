package ducksimulator

import (
	"context"
	"time"
)

func Execute(ctx context.Context, duckNameList []string) error {
	for _, duckName := range duckNameList {
		duck, err := createDuck(ctx, duckName)
		if err != nil {
			return err
		}

		time.Sleep(2 * time.Second)
		duckSimulator := NewDuckSimulator()
		if err := duckSimulator.Simulate(ctx, duck); err != nil {
			return err
		}
	}

	return nil
}

func createDuck(ctx context.Context, duckName string) (IDuck, error) {
	var duck IDuck
	switch duckName {
	case "MalladDuck":
		duck = NewMalladDuck()
	case "RubberDuck":
		duck = NewRubberDuck()
	case "ExtremeDuck":
		duck = NewExtremeDuck()
	}

	return duck, nil
}
