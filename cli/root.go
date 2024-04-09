package cli

import (
	design_patter_go "github.com/Torafugu2929/design_pattern_go"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "design-pattern-go",
	Short:   "Design patter implementation on Go",
	Version: design_patter_go.VERSION,
}

func init() {
	RootCmd.AddCommand(MockCmd)
	RootCmd.AddCommand(DuckSimulatorCmd)
}
