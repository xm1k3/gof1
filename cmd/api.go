/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/xm1k3/gof1/config"
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

		db, err := config.ConnectSqlite3(databaseFlag)
		if err != nil {
			log.Fatal(err)
		}

		newController := pkg.NewController(db)

		router := gin.New()
		setupRouter(router, newController)

		router.Run(":" + port)
	},
}

func setupRouter(router *gin.Engine, controller pkg.Controller) {
	v1 := router.Group("v1")
	v1.GET("/driver/:id", api.GetDriver(controller))
	v1.GET("/drivers/", api.GetDrivers(controller))
	v1.GET(":year/drivers/", api.GetDriversByYear(controller))
}

func init() {
	rootCmd.AddCommand(apiCmd)
	apiCmd.Flags().StringP("port", "p", "8080", "port for api")
}
