package trecker

import (
	"github.com/mslacken/treckerq/internal/pkg/help"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:           "trecker COMMAND [OPTIONS] list",
		Short:         "Executes command given in list file",
		Long:          `Executes the commands given in a list`,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	verboseArg bool
	DebugFlag  bool
)

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verboseArg, "verbose", "v", false, "Run with increased verbosity.")
	rootCmd.PersistentFlags().BoolVarP(&DebugFlag, "debug", "d", false, "Run with debugging messages enabled.")

	rootCmd.SetUsageTemplate(help.UsageTemplate)
	rootCmd.SetHelpTemplate(help.HelpTemplate)
	viper.AutomaticEnv()

}
