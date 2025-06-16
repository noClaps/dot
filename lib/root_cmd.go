package lib

import (
	"github.com/pol-rivero/doot/lib/install"
	"github.com/spf13/cobra"
)

func ExecuteRootCmd(cmd *cobra.Command, rawArgs []string) {
	ExecuteInstall(cmd)
}

func ExecuteInstall(cmd *cobra.Command) {
	fullClean, err := cmd.Flags().GetBool("full-clean")
	if err != nil {
		panic(err)
	}
	install.Install(fullClean)
}
