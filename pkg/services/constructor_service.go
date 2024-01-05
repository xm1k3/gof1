package services

import (
	"strconv"

	"github.com/xm1k3/gof1/pkg/models"
)

func (f F1Service) GetConstructor(id int) (models.Constructor, error) {
	return f.Repository.GetConstructor(id)
}

func (f F1Service) GetConstructorsByYear(year int) ([]models.Constructor, error) {
	return f.Repository.GetConstructorsByYear(year)
}

func (f F1Service) ImportConstructorsFromCsv(record []string) error {
	constructorID, err := strconv.Atoi(record[0])
	if err != nil {
		return err
	}

	constructor := models.Constructor{
		ConstructorID:  constructorID,
		ConstructorRef: record[1],
		Name:           record[2],
		Nationality:    record[3],
		URL:            record[4],
	}

	f.Repository.ImportConstructorsFromCsv(constructor)
	return nil
}
