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

func init() {
	rootCmd.AddGroup(basicCommandsGroup)

	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
}
