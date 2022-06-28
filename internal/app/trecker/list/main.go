package list

import (
	"github.com/mslacken/treckerq/internal/pkg/tlog"
	"github.com/spf13/cobra"
)

func CobraRunE(cmd *cobra.Command, args []string) error {
	tlog.Info("Called list")

	return nil
}
