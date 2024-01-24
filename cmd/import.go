/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/xm1k3/gof1/config"
	"github.com/xm1k3/gof1/pkg"
	"github.com/xm1k3/gof1/pkg/models"
	"gorm.io/gorm"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import data from CSV",
	Long:  `Import data from CSV`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := config.ConnectSqlite3("f1db.db")
		if err != nil {
			log.Fatal(err)
		}

		opts := pkg.Options{
			Database: "f1db.db",
		}

		newController := pkg.NewController(opts)
		newController.DB.AutoMigrate(&models.Driver{}, &models.Circuit{}, &models.Race{}, &models.Constructor{}, &models.Result{})

		err = importCSVData(newController.DB, "data/drivers.csv", newController.Service.ImportDriversFromCsv)
		if err != nil {
			log.Fatal(err)
		}
		err = importCSVData(db, "data/constructors.csv", newController.Service.ImportConstructorsFromCsv)
		if err != nil {
			log.Fatal(err)
		}
		err = importCSVData(db, "data/circuits.csv", newController.Service.ImportCircuitsFromCsv)
		if err != nil {
			log.Fatal(err)
		}
		err = importCSVData(db, "data/races.csv", newController.Service.ImportRacesFromCsv)
		if err != nil {
			log.Fatal(err)
		}

		err = importCSVData(db, "data/results.csv", newController.Service.ImportResultFromCsv)
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(importCmd)
}

func importCSVData(db *gorm.DB, filePath string, importFunc func([]string) error) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records[1:] {
		err := importFunc(record)
		if err != nil {
			return err
		}
	}
	return nil
}
