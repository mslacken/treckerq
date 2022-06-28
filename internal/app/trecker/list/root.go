package list

import (
	"github.com/spf13/cobra"
)

var (
	baseCmd = &cobra.Command{
		Use:   "list LISTFILE",
		Short: "list all jobs in file",
		Long:  `List all jobs in the queue file`,
		RunE:  CobraRunE,
		Args:  cobra.MinimumNArgs(1),
	}
	listFile string
)

func init() {
}

// GetRootCommand returns the root cobra.Command for the application.
func GetCommand() *cobra.Command {
	return baseCmd
}
