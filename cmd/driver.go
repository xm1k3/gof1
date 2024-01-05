/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/xm1k3/gof1/config"
	"github.com/xm1k3/gof1/pkg"
)

// driverCmd represents the driver command
var driverCmd = &cobra.Command{
	Use:   "driver",
	Short: "driver command",
	Long:  `driver command`,
	Run: func(cmd *cobra.Command, args []string) {
		databaseFlag, _ := rootCmd.PersistentFlags().GetString("database")
		yearFlag, _ := cmd.Flags().GetInt("year")
		db, err := config.ConnectSqlite3(databaseFlag)
		if err != nil {
			log.Fatal(err)
		}

		newController := pkg.NewController(db)
		drivers, err := newController.Service.GetDriversByYear(yearFlag)
		if err != nil {
			log.Fatal(err)
		}
		for i, driver := range drivers {
			fmt.Println(i+1, driver.Surname, driver.Forename)
		}
	},
}

var driverGetCmd = &cobra.Command{
	Use:   "get",
	Short: "driver get command",
	Long:  `driver get command`,
	Run: func(cmd *cobra.Command, args []string) {
		databaseFlag, _ := rootCmd.PersistentFlags().GetString("database")
		driverIdFlag, _ := cmd.Flags().GetInt("id")
		db, err := config.ConnectSqlite3(databaseFlag)
		if err != nil {
			log.Fatal(err)
		}

		newController := pkg.NewController(db)
		driver, err := newController.Service.GetDriver(driverIdFlag)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", driver)
	},
}

func init() {
	rootCmd.AddCommand(driverCmd)
	driverCmd.AddCommand(driverGetCmd)

	driverCmd.Flags().IntP("year", "y", time.Now().Year(), "driver ID")
	driverGetCmd.Flags().IntP("id", "", 0, "driver ID")
}
