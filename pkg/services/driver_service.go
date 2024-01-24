package services

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/xm1k3/gof1/pkg/models"
)

func (f F1Service) AddDriver(driver models.Driver) error {
	if driver.DriverRef == "" || driver.Forename == "" || driver.Surname == "" || driver.Nationality == "" {
		return errors.New("missing required driver information")
	}

	if !driver.DOB.IsZero() && driver.DOB.After(time.Now()) {
		return errors.New("driver's date of birth cannot be in the future")
	}

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

func (f F1Service) GetDriverStandingsByYear(year int) ([]models.DriverStanding, error) {
	if year < 1950 || year > time.Now().Year() {
		return nil, fmt.Errorf("year is out of valid range")
	}

	standings, err := f.Repository.GetDriverStandingsByYear(year)
	if err != nil {
		return nil, fmt.Errorf("error retrieving driver standings: %w", err)
	}

	if len(standings) == 0 {
		return nil, fmt.Errorf("no driver standings found for year %d", year)
	}

	return standings, nil
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
