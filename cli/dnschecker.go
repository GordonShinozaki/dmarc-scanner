package cli

import (
	"errors"
	"fmt"
	"main/cli/service"

	"github.com/spf13/cobra"
)

var dnsCheckerCmd = &cobra.Command{
	Use:   "check-dns",
	Short: "This command checks dns",
	Long:  `full list of dns`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < DESIRED_ARGS {
			err := errors.New("requires a string input")
			fmt.Fprintln(Out, err)
			return err
		} else if len(args) > DESIRED_ARGS {
			err := errors.New("too many strings. Please use only one string")
			fmt.Fprintln(Out, err)
			return err
		}
		return nil
	},
	Run: service.DnsService,
}

func init() {
	RootCmd.AddCommand(dnsCheckerCmd)
}
