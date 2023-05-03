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

	// Country Routes
	r.GET("/countries", controller.ListCountries)
	r.GET("/country/:id", controller.GetCountry)
	r.POST("/countries", controller.CreateCountry)
	r.PATCH("/countries", controller.UpdateCountry)
	r.DELETE("/countries/:id", controller.DeleteCountry)

	// Province Routes
	r.GET("/provinces", controller.ListProvinces)
	r.GET("/province/:id", controller.GetProvince)
	r.POST("/provinces", controller.CreateProvince)
	r.PATCH("/provinces", controller.UpdateProvince)
	r.DELETE("/provinces/:id", controller.DeleteProvince)

	// Plant Routes
	r.GET("/plants", controller.ListPlants)
	r.GET("/plant/:id", controller.GetPlant)
	r.POST("/plants", controller.CreatePlant)
	r.PATCH("/plants", controller.UpdatePlant)
	r.DELETE("/plants/:id", controller.DeletePlant)

	// Process Routes
	r.GET("/processes", controller.ListProcesses)
	r.GET("/process/:id", controller.GetProcess)
	r.POST("/processes", controller.CreateProcess)
	r.PATCH("/processes", controller.UpdateProcess)
	r.DELETE("/processes/:id", controller.DeleteProcess)

	// ProcessLine Routes
	r.GET("/processlines", controller.ListProcessLines)
	r.GET("/processline/:id", controller.GetProcessLine)
	r.POST("/processlines", controller.CreateProcessLine)
	r.PATCH("/processlines", controller.UpdateProcessLine)
	r.DELETE("/processlines/:id", controller.DeleteProcessLine)

	// Run the server
	r.Run()
}
