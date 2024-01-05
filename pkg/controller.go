package pkg

import (
	"github.com/xm1k3/gof1/pkg/models"
	"github.com/xm1k3/gof1/pkg/repositories"
	"github.com/xm1k3/gof1/pkg/services"
	"gorm.io/gorm"
)

type Controller struct {
	Service F1Service
	DB      *gorm.DB
}

func NewController(db *gorm.DB) Controller {
	repositories := repositories.F1Repository{
		DB: db,
	}
	service := services.F1Service{
		Repository: repositories,
	}
	c := Controller{
		Service: service,
		DB:      db,
	}
	return c
}

type F1Service interface {
	ImportDriversFromCsv(record []string) error
	ImportCircuitsFromCsv(record []string) error
	ImportConstructorsFromCsv(record []string) error
	ImportRacesFromCsv(record []string) error
}

type F1Repository interface {
	ImportDrives(driver models.Driver)
	ImportCircuitsFromCsv(circuit models.Circuit)
	ImportConstructorsFromCsv(constructor models.Constructor)
	ImportRacesFromCsv(race models.Race)
}
