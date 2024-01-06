package repositories

import "github.com/xm1k3/gof1/pkg/models"

func (f F1Repository) AddResult(result models.Result) error {
	return f.DB.Create(&result).Error
}

func (f F1Repository) GetResult(id int) (models.Result, error) {
	var result models.Result
	err := f.DB.First(&result, id).Error
	return result, err
}

func (f F1Repository) UpdateResult(result models.Result) error {
	return f.DB.Save(&result).Error
}

func (f F1Repository) DeleteResult(id int) error {
	return f.DB.Delete(&models.Result{}, id).Error
}

func (f F1Repository) GetAllResults(page int, limit int) ([]models.Result, error) {
	var results []models.Result
	offset := (page - 1) * limit
	err := f.DB.Offset(offset).Limit(limit).Find(&results).Error
	return results, err
}

func (f F1Repository) ImportResultFromCsv(result models.Result) {
	f.DB.Create(&result)
}
