package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/watcharaphol1938/Web-App/entity"
)

// POST /autonomousmobilerobots
func CreateAutonomousMobileRobot(c *gin.Context) {
	var autonomousmobilerobot entity.AutonomousMobileRobot
	if err := c.ShouldBindJSON(&autonomousmobilerobot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&autonomousmobilerobot).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": autonomousmobilerobot})
}

// GET /autonomousmobilerobot/:id
func GetAutonomousMobileRobot(c *gin.Context) {
	var autonomousmobilerobot entity.AutonomousMobileRobot
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM autonomous_mobile_robots WHERE id = ?", id).Scan(&autonomousmobilerobot).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": autonomousmobilerobot})
}

// GET /autonomousmobilerobots
func ListAutonomousMobileRobots(c *gin.Context) {
	var autonomousmobilerobots []entity.AutonomousMobileRobot
	if err := entity.DB().Raw("SELECT * FROM autonomous_mobile_robots").Scan(&autonomousmobilerobots).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": autonomousmobilerobots})
}

// DELETE /autonomousmobilerobots/:id
func DeleteAutonomousMobileRobot(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM autonomous_mobile_robots WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "autonomousmobilerobot not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /autonomousmobilerobots
func UpdateAutonomousMobileRobot(c *gin.Context) {
	var autonomousmobilerobot entity.AutonomousMobileRobot
	if err := c.ShouldBindJSON(&autonomousmobilerobot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", autonomousmobilerobot.ID).First(&autonomousmobilerobot); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "autonomousmobilerobot not found"})
		return
	}

	if err := entity.DB().Save(&autonomousmobilerobot).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": autonomousmobilerobot})
}
