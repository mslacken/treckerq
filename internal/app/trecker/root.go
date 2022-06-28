package trecker

import (
	"github.com/mslacken/treckerq/internal/app/trecker/list"
	"github.com/mslacken/treckerq/internal/pkg/help"
	"github.com/mslacken/treckerq/internal/pkg/tlog"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "trecker COMMAND [OPTIONS] LISTFILE",
		Short:                 "Executes command given in list file",
		Long:                  `Control for the trecker command`,
		PersistentPreRunE:     rootPersistentPreRunE,
		SilenceUsage:          true,
		SilenceErrors:         true,
	}
	verboseArg bool
	DebugFlag  bool
)

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// GetRootCommand returns the root cobra.Command for the application.
func GetRootCommand() *cobra.Command {
	return rootCmd
}

func init() {
	viper.AutomaticEnv()

	rootCmd.PersistentFlags().BoolVarP(&verboseArg, "verbose", "v", false, "Run with increased verbosity.")
	rootCmd.PersistentFlags().BoolVarP(&DebugFlag, "debug", "d", false, "Run with debugging messages enabled.")

	rootCmd.SetUsageTemplate(help.UsageTemplate)
	rootCmd.SetHelpTemplate(help.HelpTemplate)

	rootCmd.AddCommand(list.GetCommand())
}
func rootPersistentPreRunE(cmd *cobra.Command, args []string) error {
	if DebugFlag {
		tlog.SetLogLevel(tlog.DEBUG)
	} else if verboseArg {
		tlog.SetLogLevel(tlog.VERBOSE)
	} else {
		tlog.SetLogLevel(tlog.INFO)
	}
	return nil
}
