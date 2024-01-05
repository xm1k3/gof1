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
	AddDriver(driver models.Driver) error
	GetDriver(id int) (models.Driver, error)
	GetDrivers(page, limit int) ([]models.Driver, error)
	GetDriversByYear(year int) ([]models.Driver, error)
	UpdateDriver(driver models.Driver) error
	DeleteDriver(id int) error

	GetConstructor(id int) (models.Constructor, error)
	GetConstructorsByYear(year int) ([]models.Constructor, error)

	ImportDriversFromCsv(record []string) error
	ImportCircuitsFromCsv(record []string) error
	ImportConstructorsFromCsv(record []string) error
	ImportRacesFromCsv(record []string) error
	ImportResultFromCsv(record []string) error
}

type F1Repository interface {
	AddDriver(driver models.Driver) error
	GetDriver(id int) (models.Driver, error)
	GetDrivers(page int, limit int) ([]models.Driver, error)
	GetDriversByYear(year int) ([]models.Driver, error)
	UpdateDriver(driver models.Driver) error
	DeleteDriver(id int) error

	AddResult(result models.Result) error
	GetResult(id int) (models.Result, error)
	UpdateResult(result models.Result) error
	DeleteResult(id int) error

	GetConstructor(id int) (models.Constructor, error)
	GetConstructorsByYear(year int) ([]models.Constructor, error)

	ImportDrives(driver models.Driver)
	ImportCircuitsFromCsv(circuit models.Circuit)
	ImportConstructorsFromCsv(constructor models.Constructor)
	ImportRacesFromCsv(race models.Race)
	ImportResultFromCsv(result models.Result)
}
