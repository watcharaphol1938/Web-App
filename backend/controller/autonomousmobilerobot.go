package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/watcharaphol1938/Web-App/entity"
)

// POST /autonomousmobilerobots
func CreateAutonomousMobileRobot(c *gin.Context) {
	var autonomousmobilerobot entity.AutonomousMobileRobot
	var country entity.Country
	var province entity.Province
	var plant entity.Plant
	var process entity.Process
	var processline entity.ProcessLine
	if err := c.ShouldBindJSON(&autonomousmobilerobot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Table("countries").Where("id = ?", autonomousmobilerobot.ProcessLine.Process.Plant.Province.CountryID).First(&country); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "country not found"})
		return
	}
	if tx := entity.DB().Table("provinces").Where("id = ?", autonomousmobilerobot.ProcessLine.Process.Plant.ProvinceID).First(&province); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "province not found"})
		return
	}
	if tx := entity.DB().Table("plant").Where("id = ?", autonomousmobilerobot.ProcessLine.Process.PlantID).First(&plant); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "plant not found"})
		return
	}
	if tx := entity.DB().Table("process").Where("id = ?", autonomousmobilerobot.ProcessLine.ProcessID).First(&process); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "process not found"})
		return
	}
	if tx := entity.DB().Table("processline").Where("id = ?", autonomousmobilerobot.ProcessLineID).First(&processline); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "processline not found"})
		return
	}
	amr := entity.AutonomousMobileRobot{
		Name:                    autonomousmobilerobot.Name,
		SingleBoardComputerName: autonomousmobilerobot.SingleBoardComputerName,
		MicrocontrollerName:     autonomousmobilerobot.MicrocontrollerName,
		LiDARName:               autonomousmobilerobot.LiDARName,
		MotorName:               autonomousmobilerobot.MotorName,
		MotorAmount:             autonomousmobilerobot.MotorAmount,
		WheelName:               autonomousmobilerobot.WheelName,
		WheelAmount:             autonomousmobilerobot.WheelAmount,
		EncoderName:             autonomousmobilerobot.EncoderName,
		MotorDriverName:         autonomousmobilerobot.MotorDriverName,
		BatteryName:             autonomousmobilerobot.BatteryName,
		CameraName:              autonomousmobilerobot.CameraName,
		Country:                 country,
		Province:                province,
		Plant:                   plant,
		Process:                 process,
		ProcessLine:             processline,
	}
	if err := entity.DB().Create(&amr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": autonomousmobilerobot})
}

// GET /autonomousmobilerobot/:id
func GetAutonomousMobileRobot(c *gin.Context) {
	var autonomousmobilerobot entity.AutonomousMobileRobot
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM autonomous_mobile_robots WHERE id = ?", id).
		Preload("Country").
		Preload("Province").
		Preload("Plant").
		Preload("Process").
		Preload("ProcessLine").
		Find(&autonomousmobilerobot).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": autonomousmobilerobot})
}

// GET /autonomousmobilerobots
func ListAutonomousMobileRobots(c *gin.Context) {
	var autonomousmobilerobots []entity.AutonomousMobileRobot
	if err := entity.DB().Raw("SELECT * FROM autonomous_mobile_robots").
		Preload("Country").
		Preload("Province").
		Preload("Plant").
		Preload("Process").
		Preload("ProcessLine").
		Find(&autonomousmobilerobots).Error; err != nil {
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
	var payload entity.AutonomousMobileRobot
	var country entity.Country
	var province entity.Province
	var plant entity.Plant
	var process entity.Process
	var processline entity.ProcessLine
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", payload.ID).First(&autonomousmobilerobot); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "autonomousmobilerobot_id not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", payload.CountryID).First(&country); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "country_id not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", payload.ProvinceID).First(&province); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "province_id not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", payload.PlantID).First(&plant); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "plant_id not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", payload.ProcessID).First(&process); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "process_id not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", payload.ProcessLineID).First(&processline); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "processline_id not found"})
		return
	}
	updateamr := entity.AutonomousMobileRobot{
		Name:                    payload.Name,
		SingleBoardComputerName: payload.SingleBoardComputerName,
		MicrocontrollerName:     payload.MicrocontrollerName,
		LiDARName:               payload.LiDARName,
		MotorName:               payload.MotorName,
		MotorAmount:             payload.MotorAmount,
		WheelName:               payload.WheelName,
		WheelAmount:             payload.WheelAmount,
		EncoderName:             payload.EncoderName,
		MotorDriverName:         payload.MotorDriverName,
		BatteryName:             payload.BatteryName,
		CameraName:              payload.CameraName,
		Country:                 country,
		Province:                province,
		Plant:                   plant,
		Process:                 process,
		ProcessLine:             processline,
	}
	if err := entity.DB().Where("id = ?", autonomousmobilerobot.ID).Updates(&updateamr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// if tx := entity.DB().Where("id = ?", autonomousmobilerobot.ID).First(&autonomousmobilerobot); tx.RowsAffected == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "autonomousmobilerobot not found"})
	// 	return
	// }

	// if err := entity.DB().Save(&autonomousmobilerobot).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"data": autonomousmobilerobot})
}
