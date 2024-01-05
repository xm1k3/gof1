package repositories

import "github.com/xm1k3/gof1/pkg/models"

func (f F1Repository) ImportRacesFromCsv(race models.Race) {
	f.DB.Create(&race)
}
