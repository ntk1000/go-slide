package crawler

import (
	"github.com/spf13/cobra"
	// any imports
)

// New cobra command
func New() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "a",
		Short: "xxx",
		Long:  `xxx batch.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Docker process
			for {
				// any codes
				// e.x. polling to AWS SQS
			}
		},
	}

	return cmd
}
