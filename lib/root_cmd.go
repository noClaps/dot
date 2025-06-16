package lib

import (
	"github.com/noclaps/dot/lib/install"
	"github.com/spf13/cobra"
)

func ExecuteRootCmd(cmd *cobra.Command) {
	fullClean, err := cmd.Flags().GetBool("full-clean")
	if err != nil {
		panic(err)
	}
	install.Install(fullClean)
}
