package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var MockCmd = &cobra.Command{
	Use:   "mock",
	Short: "Mock of the command",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Hello design pattern on Go!")
		return nil
	},
}
