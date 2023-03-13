package cli

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

const DESIRED_ARGS int = 1

var Out io.Writer = os.Stdout

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "dmarc-scanner",
	Short: "A small cli based dns checker tool.",
	Long:  `A small cli based dns checker tool.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}
