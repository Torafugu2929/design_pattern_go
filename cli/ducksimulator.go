package cli

import (
	"github.com/Torafugu2929/design_pattern_go/src/ducksimulator"
	"github.com/spf13/cobra"
)

var DuckSimulatorCmd = &cobra.Command{
	Use:   "duck-simulator",
	Short: "Duck simulator by Strategy pattern",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ducksimulator.Execute(cmd.Context(), args)
		if err != nil {
			return err
		}
		return nil
	},
}
