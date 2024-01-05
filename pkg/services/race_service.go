package services

import (
	"strconv"
	"time"

	"github.com/xm1k3/gof1/pkg/models"
)

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
