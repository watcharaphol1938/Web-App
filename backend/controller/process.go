package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/watcharaphol1938/Web-App/entity"
)

// POST /processes
func CreateProcess(c *gin.Context) {
	var process entity.Process
	if err := c.ShouldBindJSON(&process); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&process).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": process})
}

// GET /process/:id
func GetProcess(c *gin.Context) {
	var process entity.Process
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM processes WHERE id = ?", id).Scan(&process).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": process})
}

// GET /processes
func ListProcesses(c *gin.Context) {
	var processes []entity.Process
	if err := entity.DB().Raw("SELECT * FROM processes").Scan(&processes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": processes})
}

// DELETE /processes/:id
func DeleteProcess(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM processes WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "process not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /processes
func UpdateProcess(c *gin.Context) {
	var process entity.Process
	if err := c.ShouldBindJSON(&process); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", process.ID).First(&process); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "process not found"})
		return
	}

	if err := entity.DB().Save(&process).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": process})
}
