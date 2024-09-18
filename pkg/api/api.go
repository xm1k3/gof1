package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xm1k3/gof1/pkg"
	"github.com/xm1k3/gof1/pkg/models"
)

// BasicAuth middleware
// Username: admin, Password: password
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"admin": "password",
	})
}

func SetupRouter(router *gin.Engine, controller pkg.Controller) {
	v1 := router.Group("/v1")
	{
		v1.GET("/driver/:id", GetDriver(controller))
		v1.GET("/drivers/", GetDrivers(controller))
		v1.GET("/drivers/year/:year", GetDriversByYear(controller))
		v1.GET("/drivers/standings/:year", GetDriverStandingsByYear(controller))
	}

	v1Auth := router.Group("/v1")
	{
		v1Auth.Use(BasicAuth())
		{
			v1Auth.POST("/drivers", AddDriver(controller))
		}
	}
}

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
		driver, err := controller.Service.GetDriversByYear(driverID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, driver)
	}
}

func GetDriverStandingsByYear(controller pkg.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {
		driverIDStr := c.Param("year")
		driverID, err := strconv.Atoi(driverIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year param"})
			return
		}
		driverStanding, err := controller.Service.GetDriverStandingsByYear(driverID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, driverStanding)
	}
}

func AddDriver(controller pkg.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newDriver models.Driver
		if err := c.ShouldBindJSON(&newDriver); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := controller.Service.AddDriver(newDriver); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Driver added successfully"})
	}
}
