package services

import (
	"strconv"

	"github.com/xm1k3/gof1/pkg/models"
)

func (f F1Service) AddCircuit(circuit models.Circuit) error {
	return f.Repository.AddCircuit(circuit)
}

func (f F1Service) GetCircuit(id int) (models.Circuit, error) {
	return f.Repository.GetCircuit(id)
}

func (f F1Service) GetCircuits(page, limit int) ([]models.Circuit, error) {
	return f.Repository.GetCircuits(page, limit)
}

func (f F1Service) GetCircuitsByYear(year int) ([]models.Circuit, error) {
	return f.Repository.GetCircuitsByYear(year)
}

func (f F1Service) UpdateCircuit(circuit models.Circuit) error {
	return f.Repository.UpdateCircuit(circuit)
}

func (f F1Service) DeleteCircuit(id int) error {
	return f.Repository.DeleteCircuit(id)
}

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
