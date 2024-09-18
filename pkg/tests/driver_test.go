package tests

import (
	"encoding/csv"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/xm1k3/gof1/pkg"
	"github.com/xm1k3/gof1/pkg/api"
	"gorm.io/gorm"
)

func TestGetDriver(t *testing.T) {
	// Setup
	controller := pkg.NewController(pkg.Options{
		Database: "test.db",
	})

	err := importCSVData(controller.DB, "../../data/drivers.csv", controller.Service.ImportDriversFromCsv)
	if err != nil {
		log.Fatal(err)
	}

	// Test case
	t.Run("GetDriver - Valid ID", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/v1/driver/1", nil)
		resp := httptest.NewRecorder()

		router := gin.New()
		api.SetupRouter(router, controller)
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Result())
		// Add more assertions as needed
	})
}

func importCSVData(db *gorm.DB, filePath string, importFunc func([]string) error) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records[1:] {
		err := importFunc(record)
		if err != nil {
			return err
		}
	}
	return nil
}
