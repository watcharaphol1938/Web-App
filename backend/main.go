package main

import (
	"github.com/gin-gonic/gin"
	"github.com/watcharaphol1938/Web-App/controller"
	"github.com/watcharaphol1938/Web-App/entity"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()

	// AutonomousMobileRobot Routes
	r.GET("/autonomousmobilerobots", controller.ListAutonomousMobileRobots)
	r.GET("/autonomousmobilerobot/:id", controller.GetAutonomousMobileRobot)
	r.POST("/autonomousmobilerobots", controller.CreateAutonomousMobileRobot)
	r.PATCH("/autonomousmobilerobots", controller.UpdateAutonomousMobileRobot)
	r.DELETE("/autonomousmobilerobots/:id", controller.DeleteAutonomousMobileRobot)

	// Run the server
	r.Run()
}
