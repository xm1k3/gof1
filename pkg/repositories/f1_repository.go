package repositories

import (
	"github.com/xm1k3/gof1/pkg/models"
	"gorm.io/gorm"
)

type F1Repository struct {
	DB *gorm.DB
}

func (f F1Repository) ImportDriversFromCsv(driver models.Driver) {
	f.DB.Create(&driver)
}

func (f F1Repository) ImportConstructorsFromCsv(constructor models.Constructor) {
	f.DB.Create(&constructor)
}

func (f F1Repository) ImportCircuitsFromCsv(circuit models.Circuit) {
	f.DB.Create(&circuit)
}

func (f F1Repository) ImportRacesFromCsv(race models.Race) {
	f.DB.Create(&race)
}
