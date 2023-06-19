package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"otus-recipe/app/config"
	"otus-recipe/app/server"
)

var apiServer = &cobra.Command{
	Use:   "application",
	Short: "Run api http server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.NewConfig()

		if err != nil {
			log.Fatal(err)
		}

		s := server.New(config)

		if err := s.Start(); err != nil {
			log.Fatal(err)
		}
	},
}
