package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"otus-recipe/app/server"
)

var apiServer = &cobra.Command{
	Use:   "application",
	Short: "Run api http server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		s := server.New(container.Config, container.Log, container.Processors)

		if err := s.Start(); err != nil {
			log.Fatal(err)
		}
	},
}
