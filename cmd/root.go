package cmd

import (
	"os"

	"github.com/noclaps/dot/lib"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dot",
	Short: "A fast and simple dotfiles manager that just gets the job done.",
	Run: func(cmd *cobra.Command, args []string) {
		lib.ExecuteRootCmd(cmd)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var basicCommandsGroup = &cobra.Group{
	ID:    "basicCommands",
	Title: "Basic commands:",
}

var otherCommandsGroup = &cobra.Group{
	ID:    "otherCommands",
	Title: "Other commands:",
}

func init() {
	rootCmd.AddGroup(basicCommandsGroup)
	rootCmd.AddGroup(otherCommandsGroup)

	rootCmd.Flags().Bool("full-clean", false, "Search and remove all broken symlinks that point to the dotfiles directory, even if they were created by another program. Can be slow.")

	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.SetHelpCommandGroupID(otherCommandsGroup.ID)
	rootCmd.SetCompletionCommandGroupID(otherCommandsGroup.ID)
}
