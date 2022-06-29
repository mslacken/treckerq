package count

import (
	"github.com/spf13/cobra"
)

var (
	baseCmd = &cobra.Command{
		Use:   "count LISTFILE",
		Short: "print information about the queue",
		Long:  `List all jobs in the queue file`,
		RunE:  CobraRunE,
		Args:  cobra.MinimumNArgs(1),
	}
	QueuedTasks bool
)

func init() {
	baseCmd.PersistentFlags().BoolVar(&QueuedTasks, "Q", false, "print number of tasks left in queue")
}

// GetRootCommand returns the root cobra.Command for the application.
func GetCommand() *cobra.Command {
	return baseCmd
}
