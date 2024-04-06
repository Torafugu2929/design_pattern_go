package main

import (
	"context"

	"github.com/Torafugu2929/design_pattern_go/cli"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := cli.RootCmd.ExecuteContext(ctx); err != nil {
		panic(err)
	}
}
