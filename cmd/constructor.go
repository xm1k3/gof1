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

// constructorCmd represents the constructor command
var constructorCmd = &cobra.Command{
	Use:   "constructor",
	Short: "constructor command",
	Long:  `constructor command`,
	Run: func(cmd *cobra.Command, args []string) {
		databaseFlag, _ := rootCmd.PersistentFlags().GetString("database")
		yearFlag, _ := cmd.Flags().GetInt("year")
		db, err := config.ConnectSqlite3(databaseFlag)
		if err != nil {
			log.Fatal(err)
		}

		newController := pkg.NewController(db)
		constructors, err := newController.Service.GetConstructorsByYear(yearFlag)
		if err != nil {
			log.Fatal(err)
		}
		for i, constructor := range constructors {
			fmt.Println(i+1, constructor.Name)
		}
	},
}

var constructorGetCmd = &cobra.Command{
	Use:   "get",
	Short: "constructor get command",
	Long:  `constructor get command`,
	Run: func(cmd *cobra.Command, args []string) {
		databaseFlag, _ := rootCmd.PersistentFlags().GetString("database")
		constructorIdFlag, _ := cmd.Flags().GetInt("id")
		db, err := config.ConnectSqlite3(databaseFlag)
		if err != nil {
			log.Fatal(err)
		}

		newController := pkg.NewController(db)
		constructor, err := newController.Service.GetDriver(constructorIdFlag)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", constructor)
	},
}

func init() {
	rootCmd.AddCommand(constructorCmd)
	constructorCmd.AddCommand(constructorGetCmd)

	constructorCmd.Flags().IntP("year", "y", time.Now().Year(), "year")
	constructorGetCmd.Flags().IntP("id", "", 0, "constructor ID")
}
