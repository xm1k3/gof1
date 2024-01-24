package repositories

import "github.com/xm1k3/gof1/pkg/models"

func (f F1Repository) GetConstructor(id int) (models.Constructor, error) {
	var constructor models.Constructor
	result := f.DB.First(&constructor, id)
	return constructor, result.Error
}

func (f F1Repository) GetConstructorsByYear(year int) ([]models.Constructor, error) {
	var constructors []models.Constructor
	err := f.DB.
		Joins("JOIN results on results.constructorId = constructors.id").
		Joins("JOIN races on races.id = results.raceId").
		Where("races.year = ?", year).
		Distinct().
		Find(&constructors).Error
	return constructors, err
}

func (f F1Repository) GetConstructorsStandingsByYear(year int) ([]models.ConstructorStanding, error) {
	var standings []models.ConstructorStanding

	err := f.DB.
		Table("results").
		Select("constructors.id, constructors.name, SUM(results.points) as points").
		Joins("JOIN constructors on constructors.id = results.constructorId").
		Joins("JOIN races on races.id = results.raceId").
		Where("races.year = ?", year).
		Group("constructors.id, constructors.name").
		Order("SUM(results.points) DESC").
		Scan(&standings).Error

	return standings, err
}

func (f F1Repository) ImportConstructorsFromCsv(constructor models.Constructor) {
	f.DB.Create(&constructor)
}
