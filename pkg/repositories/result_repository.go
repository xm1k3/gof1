package repositories

import "github.com/xm1k3/gof1/pkg/models"

func (f F1Repository) ImportResultFromCsv(result models.Result) {
	f.DB.Create(&result)
}
