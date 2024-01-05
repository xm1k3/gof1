package services

import (
	"strconv"

	"github.com/xm1k3/gof1/pkg/models"
)

func (f F1Service) ImportCircuitsFromCsv(record []string) error {
	circuitID, err := strconv.Atoi(record[0])
	if err != nil {
		return err
	}

	circuit := models.Circuit{
		CircuitID:  circuitID,
		CircuitRef: record[1],
		Name:       record[2],
		Location:   record[3],
		Country:    record[4],
		URL:        record[5],
	}

	f.Repository.ImportCircuitsFromCsv(circuit)
	return nil
}
