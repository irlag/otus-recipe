package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	elasticIndicesRefreshCmd = &cobra.Command{
		Use:   "elastic-recipe-load",
		Short: "Refresh elastic recipe index",
		Run: func(cmd *cobra.Command, args []string) {
			err := container.Processors.Command.RefreshElasticRecipeIndex(cmd.Context(), container.Log)
			if err != nil {
				log.Fatal(
					fmt.Sprintf("error elastic-indices-refresh"),
					zap.Error(err),
				)
			}
		},
	}
)
