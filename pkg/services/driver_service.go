package services

import (
	"strconv"
	"time"

	"github.com/xm1k3/gof1/pkg/models"
)

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
