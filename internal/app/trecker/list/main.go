package list

import (
	"fmt"

	"github.com/mslacken/treckerq/internal/pkg/queue"
	"github.com/mslacken/treckerq/internal/pkg/tlog"
	"github.com/spf13/cobra"
)

func CobraRunE(cmd *cobra.Command, args []string) error {
	tlog.Debug("Called list")
	queue, err := queue.OpenQueueFile(args[0])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("queue: %v\n", queue)
	return nil
}
