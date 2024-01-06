package repositories

import "github.com/xm1k3/gof1/pkg/models"

func (f F1Repository) AddCircuit(circuit models.Circuit) error {
	result := f.DB.Create(&circuit)
	return result.Error
}

func (f F1Repository) GetCircuit(id int) (models.Circuit, error) {
	var circuit models.Circuit
	result := f.DB.First(&circuit, id)
	return circuit, result.Error
}

func (f F1Repository) GetCircuits(page, limit int) ([]models.Circuit, error) {
	var circuits []models.Circuit
	offset := (page - 1) * limit
	result := f.DB.Offset(offset).Limit(limit).Find(&circuits)
	return circuits, result.Error
}

func (f F1Repository) GetCircuitsByYear(year int) ([]models.Circuit, error) {
	var circuits []models.Circuit
	err := f.DB.
		Joins("JOIN races on races.circuitId = circuits.id").
		Where("races.year = ?", year).
		Distinct().
		Find(&circuits).Error
	return circuits, err
}

func (f F1Repository) UpdateCircuit(circuit models.Circuit) error {
	result := f.DB.Save(&circuit)
	return result.Error
}

func (f F1Repository) DeleteCircuit(id int) error {
	result := f.DB.Delete(&models.Circuit{}, id)
	return result.Error
}

func (f F1Repository) ImportCircuitsFromCsv(circuit models.Circuit) {
	f.DB.Create(&circuit)
}
