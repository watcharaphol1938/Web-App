package entity

import (
	"gorm.io/gorm"
)

type AutonomousMobileRobot struct {
	gorm.Model
	Name                    string
	SingleBoardComputerName string
	MicrocontrollerName     string
	LiDARName               string
	MotorName               string
	MotorAmount             int
	WheelName               string
	WheelAmount             int
	EncoderName             string
	MotorDriverName         string
	BatteryName             string
	CameraName              string

	ProcessLineID *uint
	ProcessLine   ProcessLine `gorm:"references:id" valid:"-"`

	ProcessID *uint
	Process   Process `gorm:"references:id" valid:"-"`

	PlantID *uint
	Plant   Plant `gorm:"references:id" valid:"-"`

	ProvinceID *uint
	Province   Province `gorm:"references:id" valid:"-"`

	CountryID *uint
	Country   Country `gorm:"references:id" valid:"-"`
}
type ProcessLine struct {
	gorm.Model
	Name string

	ProcessID *uint
	Process   Process `gorm:"references:id" valid:"-"`

	AutonomousMobileRobots []AutonomousMobileRobot `gorm:"foreignKey:ProcessLineID"`
}
type Process struct {
	gorm.Model
	Name string

	PlantID *uint
	Plant   Plant `gorm:"references:id" valid:"-"`

	ProcessLines           []ProcessLine           `gorm:"foreignKey:ProcessID"`
	AutonomousMobileRobots []AutonomousMobileRobot `gorm:"foreignKey:ProcessLineID"`
}
type Plant struct {
	gorm.Model
	Name string

	ProvinceID *uint
	Province   Province `gorm:"references:id" valid:"-"`

	Processes              []Process               `gorm:"foreignKey:PlantID"`
	AutonomousMobileRobots []AutonomousMobileRobot `gorm:"foreignKey:ProcessLineID"`
}
type Province struct {
	gorm.Model
	Name string

	CountryID *uint
	Country   Country `gorm:"references:id" valid:"-"`

	Plants                 []Plant                 `gorm:"foreignKey:ProvinceID"`
	AutonomousMobileRobots []AutonomousMobileRobot `gorm:"foreignKey:ProcessLineID"`
}
type Country struct {
	gorm.Model
	Name string

	Provinces              []Province              `gorm:"foreignKey:CountryID"`
	AutonomousMobileRobots []AutonomousMobileRobot `gorm:"foreignKey:ProcessLineID"`
}
