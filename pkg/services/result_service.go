package services

import (
	"strconv"

	"github.com/xm1k3/gof1/pkg/models"
)

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
