package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/watcharaphol1938/Web-App/entity"
)

// POST /countries
func CreateCountry(c *gin.Context) {
	var country entity.Country
	if err := c.ShouldBindJSON(&country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&country).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": country})
}

// GET /country/:id
func GetCountry(c *gin.Context) {
	var country entity.Country
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM countries WHERE id = ?", id).Scan(&country).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": country})
}

// GET /countries
func ListCountries(c *gin.Context) {
	var countries []entity.Country
	if err := entity.DB().Raw("SELECT * FROM countries").Scan(&countries).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": countries})
}

// DELETE /countries/:id
func DeleteCountry(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM countries WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "country not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /countries
func UpdateCountry(c *gin.Context) {
	var country entity.Country
	if err := c.ShouldBindJSON(&country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", country.ID).First(&country); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "country not found"})
		return
	}

	if err := entity.DB().Save(&country).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": country})
}
