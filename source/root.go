package root

import (
	"github.com/spf13/cobra"
	// any imports
)

// New contained all command
func New() *cobra.Command {

	var rootCmd = &cobra.Command{Use: "xxxmicro"}

	rootCmd.AddCommand(a.New())
	rootCmd.AddCommand(b.New())

	return rootCmd
}
