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

func (f F1Service) AddDriver(driver models.Driver) error {
	return f.Repository.AddDriver(driver)
}

func (f F1Service) GetDriver(id int) (models.Driver, error) {
	return f.Repository.GetDriver(id)
}

func (f F1Service) GetDrivers(page, limit int) ([]models.Driver, error) {
	return f.Repository.GetDrivers(page, limit)
}

func (f F1Service) GetDriversByYear(year int) ([]models.Driver, error) {
	return f.Repository.GetDriversByYear(year)
}

func (f F1Service) UpdateDriver(driver models.Driver) error {
	return f.Repository.UpdateDriver(driver)
}

func (f F1Service) DeleteDriver(id int) error {
	return f.Repository.DeleteDriver(id)
}

func (f F1Service) ImportDriversFromCsv(record []string) error {
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

	f.Repository.ImportDriversFromCsv(driver)
	return nil
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

func (f F1Service) ImportRacesFromCsv(record []string) error {
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

	f.Repository.ImportRacesFromCsv(race)
	return nil
}

func (f F1Service) ImportResultFromCsv(record []string) error {
	resultID, err := strconv.Atoi(record[0])
	if err != nil {
		return err
	}

	raceID, _ := strconv.Atoi(record[1])
	driverID, _ := strconv.Atoi(record[2])
	constructorID, _ := strconv.Atoi(record[3])
	number, _ := strconv.Atoi(record[4])
	grid, _ := strconv.Atoi(record[5])
	position, _ := strconv.Atoi(record[6])
	positionOrder, _ := strconv.Atoi(record[8])
	points, _ := strconv.ParseFloat(record[9], 64)
	laps, _ := strconv.Atoi(record[10])
	milliseconds, _ := strconv.Atoi(record[12])
	fastestLap, _ := strconv.Atoi(record[13])
	rank, _ := strconv.Atoi(record[14])
	statusID, _ := strconv.Atoi(record[17])

	result := models.Result{
		ResultID:        resultID,
		RaceID:          raceID,
		DriverID:        driverID,
		ConstructorID:   constructorID,
		Number:          number,
		Grid:            grid,
		Position:        position,
		PositionText:    record[7],
		PositionOrder:   positionOrder,
		Points:          points,
		Laps:            laps,
		Time:            record[11],
		Milliseconds:    milliseconds,
		FastestLap:      fastestLap,
		Rank:            rank,
		FastestLapTime:  record[15],
		FastestLapSpeed: record[16],
		StatusID:        statusID,
	}

	f.Repository.ImportResultFromCsv(result)
	return nil
}
