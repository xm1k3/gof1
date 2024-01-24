package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/xm1k3/gof1/pkg/models"
)

func (f F1Service) GetConstructor(id int) (models.Constructor, error) {
	return f.Repository.GetConstructor(id)
}

func (f F1Service) GetConstructorsByYear(year int) ([]models.Constructor, error) {
	return f.Repository.GetConstructorsByYear(year)
}

func (f F1Service) GetConstructorsStandingsByYear(year int) ([]models.ConstructorStanding, error) {
	if year < 1950 || year > time.Now().Year() {
		return nil, fmt.Errorf("year is out of valid range")
	}

	standings, err := f.Repository.GetConstructorsStandingsByYear(year)
	if err != nil {
		return nil, fmt.Errorf("error retrieving constructor standings: %w", err)
	}

	if len(standings) == 0 {
		return nil, fmt.Errorf("no constructor standings found for year %d", year)
	}

	return standings, nil
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
