package repositories

import "github.com/xm1k3/gof1/pkg/models"

func (f F1Repository) ImportCircuitsFromCsv(circuit models.Circuit) {
	f.DB.Create(&circuit)
}
