package cmd

import (
	"github.com/noclaps/dot/lib/add"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	GroupID: basicCommandsGroup.ID,
	Use:     "add <file> [file2 ...]",
	Short:   "Move one or more files to the dotfiles directory and symlink them.",
	Run: func(cmd *cobra.Command, args []string) {
		add.Add(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Args = cobra.MinimumNArgs(1)
	addCmd.ArgAliases = []string{"file"}
}
