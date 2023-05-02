package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/watcharaphol1938/Web-App/entity"
)

// POST /processlines
func CreateProcessLine(c *gin.Context) {
	var processline entity.ProcessLine
	if err := c.ShouldBindJSON(&processline); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&processline).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": processline})
}

// GET /processline/:id
func GetProcessLine(c *gin.Context) {
	var processline entity.ProcessLine
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM process_lines WHERE id = ?", id).Scan(&processline).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": processline})
}

// GET /processlines
func ListProcessLines(c *gin.Context) {
	var processlines []entity.ProcessLine
	if err := entity.DB().Raw("SELECT * FROM process_lines").Scan(&processlines).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": processlines})
}

// DELETE /processlines/:id
func DeleteProcessLine(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM process_lines WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "processline not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /processlines
func UpdateProcessLine(c *gin.Context) {
	var processline entity.ProcessLine
	if err := c.ShouldBindJSON(&processline); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", processline.ID).First(&processline); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "processline not found"})
		return
	}

	if err := entity.DB().Save(&processline).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": processline})
}
