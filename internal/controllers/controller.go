package controllers

import (
	"net/http"
	"strconv"

	"github.com/fvmoraes/api-with-rabbitmq/internal/helpers"
	"github.com/fvmoraes/api-with-rabbitmq/internal/initializers"
	"github.com/fvmoraes/api-with-rabbitmq/internal/logs"
	"github.com/fvmoraes/api-with-rabbitmq/internal/models"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// CreateFoobar godoc
// @Summary Creating Foobar
// @Description Create Foobar in database
// @Tags foobar
// @Accept json
// @Produce json
// @Param foobar body models.Foobar true "Foobar Data"
// @Success 200 {object} models.Foobar
// @Router /foobar [post]
func CreateFoobar(c *gin.Context) {
	var foobar models.Foobar
	if err := c.ShouldBindJSON(&foobar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		logs.WriteLogFile("ERROR", err.Error())
		return
	}
	if err := helpers.ModelValidator(&foobar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		logs.WriteLogFile("WARNING", err.Error())
		return
	}
	initializers.DB.Create(&foobar)
	c.JSON(http.StatusOK, foobar)
	logs.WriteLogFile("INFO", "Successful call to endpoint: "+""+c.Request.Method+""+c.Request.RequestURI)
}

// @BasePath /api/v1
// ShownFoobar godoc
// @Summary Showning all Foobar
// @Description Shown all Foobar in database
// @Tags foobar
// @Accept json
// @Produce json
// @Success 200 {object} models.Foobar
// @Router /foobar [get]
func ShownFoobar(c *gin.Context) {
	var foobar []models.Foobar
	if err := initializers.DB.Find(&foobar); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": err.Error.Error(),
		})
		logs.WriteLogFile("ERROR", err.Error.Error())
		return
	}
	c.JSON(http.StatusOK, foobar)
	logs.WriteLogFile("INFO", "Successful call to endpoint: "+""+c.Request.Method+""+c.Request.RequestURI)
}

// @BasePath /api/v1
// ShownFoobarByParamId godoc
// @Summary Showning Foobar by ID
// @Description Shown Foobar by ID in database
// @Tags foobar
// @Accept json
// @Produce json
// @Param id path int true "Foobar ID"
// @Success 200 {object} models.Foobar
// @Router /foobar/{id} [get]
func ShownFoobarByParamId(c *gin.Context) {
	var foobar models.Foobar
	id := c.Params.ByName("id")
	if err := initializers.DB.First(&foobar, id); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": err.Error.Error(),
		})
		logs.WriteLogFile("ERROR", err.Error.Error())
		return
	}
	c.JSON(http.StatusOK, foobar)
	logs.WriteLogFile("INFO", "Successful call to endpoint: "+""+c.Request.Method+""+c.Request.RequestURI)
}

// @BasePath /api/v1
// EditFoobarByParamId godoc
// @Summary Editing Foobar by ID
// @Description Edit Foobar by ID in database
// @Tags foobar
// @Accept json
// @Produce json
// @Param id path int true "Foobar ID"
// @Param foobar body models.Foobar true "Foobar Data"
// @Success 200 {object} models.Foobar
// @Router /foobar/{id} [patch]
func EditFoobarByParamId(c *gin.Context) {
	var foobar models.Foobar
	id := c.Params.ByName("id")
	if err := initializers.DB.First(&foobar, id); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": err.Error.Error(),
		})
		logs.WriteLogFile("ERROR", err.Error.Error())
		return
	}
	if err := c.ShouldBindJSON(&foobar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		logs.WriteLogFile("ERROR", err.Error())
		return
	}
	if err := helpers.ModelValidator(&foobar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		logs.WriteLogFile("WARNING", err.Error())
		return
	}
	initializers.DB.Model(&foobar).Updates(foobar)
	c.JSON(http.StatusOK, foobar)
	logs.WriteLogFile("INFO", "Successful call to endpoint: "+""+c.Request.Method+""+c.Request.RequestURI)
}

// @BasePath /api/v1
// DeleteFoobarByParamId godoc
// @Summary Deleting Foobar by ID
// @Description Delete Foobar by ID in database
// @Tags foobar
// @Accept json
// @Produce json
// @Param id path int true "Foobar ID"
// @Success 200 {object} models.Foobar
// @Router /foobar/{id} [delete]
func DeleteFoobarByParamId(c *gin.Context) {
	var foobar models.Foobar
	id := c.Params.ByName("id")
	if err := initializers.DB.First(&foobar, id); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": err.Error.Error(),
		})
		logs.WriteLogFile("ERROR", err.Error.Error())
		return
	}
	initializers.DB.Delete(&foobar, id)
	c.JSON(http.StatusAccepted, foobar)
	logs.WriteLogFile("INFO", "Successful call to endpoint: "+""+c.Request.Method+""+c.Request.RequestURI)
}

// @BasePath /api/v1
// ShownFoobarByParamReg godoc
// @Summary Showning Foobar by Registration
// @Description Shown Foobar by Registration in database
// @Tags foobar
// @Accept json
// @Produce json
// @Param reg path uint64 true "Foobar Registration"
// @Success 200 {object} models.Foobar
// @Router /foobar/{reg} [get]
func ShownFoobarByParamReg(c *gin.Context) {
	var foobar []models.Foobar
	reg, _ := strconv.ParseUint(c.Params.ByName("reg"), 10, 64)
	if err := initializers.DB.Where(&models.Foobar{Registration: reg}).Find(&foobar); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": err.Error.Error(),
		})
		logs.WriteLogFile("ERROR", err.Error.Error())
		return
	}
	c.JSON(http.StatusOK, foobar)
	logs.WriteLogFile("INFO", "Successful call to endpoint: "+""+c.Request.Method+""+c.Request.RequestURI)
}
