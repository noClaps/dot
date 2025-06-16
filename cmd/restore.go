package cmd

import (
	"github.com/noclaps/dot/lib/restore"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	GroupID: basicCommandsGroup.ID,
	Use:     "restore <file> [file2 ...]",
	Short:   "Opposite of 'doot add'. Replace symlinks with the original files.",
	Run: func(cmd *cobra.Command, args []string) {
		restore.Restore(args)
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)

	restoreCmd.Args = cobra.MinimumNArgs(1)
	restoreCmd.ArgAliases = []string{"file"}
}
