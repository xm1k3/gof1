package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xm1k3/gof1/pkg"
)

func GetDriver(controller pkg.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {
		driverIDStr := c.Param("id")
		driverID, err := strconv.Atoi(driverIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid driver ID"})
			return
		}
		driver, err := controller.Service.GetDriver(driverID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, driver)
	}
}

func GetDrivers(controller pkg.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {
		pageStr := c.Query("page")
		limitStr := c.Query("limit")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page query param"})
			return
		}
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit query param"})
			return
		}
		drivers, err := controller.Service.GetDrivers(page, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, drivers)
	}
}

func GetDriversByYear(controller pkg.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {
		driverIDStr := c.Param("year")
		driverID, err := strconv.Atoi(driverIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year param"})
			return
		}
		driver, err := controller.Service.GetConstructorsByYear(driverID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, driver)
	}
}
