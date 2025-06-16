package cmd

import (
	"github.com/noclaps/dot/lib/install"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	GroupID: basicCommandsGroup.ID,
	Use:     "clean",
	Short:   "Remove all symlinks created by dot.",
	Run: func(cmd *cobra.Command, args []string) {
		install.Clean()
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	cleanCmd.Args = cobra.NoArgs
}
