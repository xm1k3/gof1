/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/xm1k3/gof1/pkg"
	"github.com/xm1k3/gof1/pkg/api"
)

var newController pkg.Controller

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "F1 API",
	Long:  `F1 API`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		databaseFlag, _ := rootCmd.PersistentFlags().GetString("database")

		opts := pkg.Options{
			Database: databaseFlag,
		}

		newController := pkg.NewController(opts)

		router := gin.New()
		api.SetupRouter(router, newController)

		router.Run(":" + port)
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
	apiCmd.Flags().StringP("port", "p", "8080", "port for api")
}
