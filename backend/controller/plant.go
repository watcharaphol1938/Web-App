package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/watcharaphol1938/Web-App/entity"
)

// POST /plants
func CreatePlant(c *gin.Context) {
	var plant entity.Plant
	if err := c.ShouldBindJSON(&plant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&plant).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": plant})
}

// GET /plant/:id
func GetPlant(c *gin.Context) {
	var plant entity.Plant
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM plants WHERE id = ?", id).Scan(&plant).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": plant})
}

// GET /plants
func ListPlants(c *gin.Context) {
	var plants []entity.Plant
	if err := entity.DB().Raw("SELECT * FROM plants").Scan(&plants).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": plants})
}

// DELETE /plants/:id
func DeletePlant(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM plants WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "plant not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /plants
func UpdatePlant(c *gin.Context) {
	var plant entity.Plant
	if err := c.ShouldBindJSON(&plant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", plant.ID).First(&plant); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "plant not found"})
		return
	}

	if err := entity.DB().Save(&plant).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": plant})
}
