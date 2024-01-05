package repositories

import "github.com/xm1k3/gof1/pkg/models"

func (f F1Repository) AddDriver(driver models.Driver) error {
	result := f.DB.Create(&driver)
	return result.Error
}

func (f F1Repository) GetDriver(id int) (models.Driver, error) {
	var driver models.Driver
	result := f.DB.First(&driver, id)
	return driver, result.Error
}

func (f F1Repository) GetDrivers(page int, limit int) ([]models.Driver, error) {
	var drivers []models.Driver
	offset := (page - 1) * limit
	result := f.DB.Offset(offset).Limit(limit).Find(&drivers)
	return drivers, result.Error
}

func (f F1Repository) GetDriversByYear(year int) ([]models.Driver, error) {
	var drivers []models.Driver
	err := f.DB.
		Joins("JOIN results on results.driverId = drivers.id").
		Joins("JOIN races on races.id = results.raceId").
		Where("races.year = ?", year).
		Distinct().
		Find(&drivers).Error
	return drivers, err
}

func (f F1Repository) UpdateDriver(driver models.Driver) error {
	result := f.DB.Save(&driver)
	return result.Error
}

func (f F1Repository) DeleteDriver(id int) error {
	result := f.DB.Delete(&models.Driver{}, id)
	return result.Error
}

func (f F1Repository) ImportDriversFromCsv(driver models.Driver) {
	f.DB.Create(&driver)
}
