package services

import (
	"strconv"
	"time"

	"github.com/xm1k3/gof1/pkg/models"
	"github.com/xm1k3/gof1/pkg/repositories"
)

type F1Service struct {
	Repository repositories.F1Repository
}

func (f F1Service) ImportDrivers(record []string) error {
	id, err := strconv.Atoi(record[0])
	if err != nil {
		return err
	}

	dob, err := time.Parse("2006-01-02", record[6])
	if err != nil {
		return err
	}

	driver := models.Driver{
		DriverID:    id,
		DriverRef:   record[1],
		Number:      record[2],
		Code:        record[3],
		Forename:    record[4],
		Surname:     record[5],
		DOB:         dob,
		Nationality: record[7],
		URL:         record[8],
	}

	f.Repository.ImportDrivers(driver)
	return nil
}

func (f F1Service) ImportConstructors(record []string) error {
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

	f.Repository.ImportConstructors(constructor)
	return nil
}

func (f F1Service) ImportCircuits(record []string) error {
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

	f.Repository.ImportCircuits(circuit)
	return nil
}

func (f F1Service) ImportRaces(record []string) error {
	raceID, err := strconv.Atoi(record[0])
	if err != nil {
		return err
	}

	year, err := strconv.Atoi(record[1])
	if err != nil {
		return err
	}

	round, err := strconv.Atoi(record[2])
	if err != nil {
		return err
	}

	circuitID, err := strconv.Atoi(record[3])
	if err != nil {
		return err
	}

	raceDate, err := time.Parse("2006-01-02", record[5])
	if err != nil {
		return err
	}

	qualiDate, _ := time.Parse("2006-01-02", record[8])
	sprintDate, _ := time.Parse("2006-01-02", record[11])

	race := models.Race{
		RaceID:     raceID,
		Year:       year,
		Round:      round,
		CircuitID:  circuitID,
		Name:       record[4],
		Date:       raceDate,
		Time:       record[6],
		URL:        record[7],
		QualiDate:  &qualiDate,
		QualiTime:  record[9],
		SprintDate: &sprintDate,
		SprintTime: record[12],
	}

	f.Repository.ImportRaces(race)
	return nil
}
