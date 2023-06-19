package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version = ""

var consoleCmd = &cobra.Command{
	Use:     "otus-recipe",
	Short:   "Education otus recipe app",
	Long:    ``,
	Version: Version,
}

func init() {
	consoleCmd.AddCommand(apiServer)
	consoleCmd.AddCommand(version)
	consoleCmd.AddCommand(migrateCmd)
}

func Execute() {
	if err := consoleCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
