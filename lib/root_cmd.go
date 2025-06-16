package lib

import (
	"github.com/noclaps/dot/lib/install"
	"github.com/spf13/cobra"
)

func ExecuteRootCmd(cmd *cobra.Command) {
	install.Install()
}
