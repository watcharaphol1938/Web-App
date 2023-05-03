package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("AMR.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&AutonomousMobileRobot{},
		&ProcessLine{},
		&Process{},
		&Plant{},
		&Province{},
		&Country{},
	)

	db = database

	// ProcessLine -----------------------------------------------------------------------------------------------------------------------------------------
	processline1 := ProcessLine{Name: "A"}
	db.Model(&ProcessLine{}).Create(&processline1)

	processline2 := ProcessLine{Name: "B"}
	db.Model(&ProcessLine{}).Create(&processline2)

	// Process -----------------------------------------------------------------------------------------------------------------------------------------
	process1 := Process{Name: "air conditioner hoses and tubes"}
	db.Model(&Process{}).Create(&process1)

	process2 := Process{Name: "relays and flashers"}
	db.Model(&Process{}).Create(&process2)

	process3 := Process{Name: "electrical components"}
	db.Model(&Process{}).Create(&process3)

	process4 := Process{Name: "air-conditioners"}
	db.Model(&Process{}).Create(&process4)

	process5 := Process{Name: "magnetos"}
	db.Model(&Process{}).Create(&process5)

	process6 := Process{Name: "audio-visual"}
	db.Model(&Process{}).Create(&process6)

	process7 := Process{Name: "automotive electronics"}
	db.Model(&Process{}).Create(&process7)

	process8 := Process{Name: "mold & die"}
	db.Model(&Process{}).Create(&process8)

	process9 := Process{Name: "spare parts"}
	db.Model(&Process{}).Create(&process9)

	process10 := Process{Name: "jig & tools"}
	db.Model(&Process{}).Create(&process10)

	process11 := Process{Name: "fuel injection system"}
	db.Model(&Process{}).Create(&process11)

	process12 := Process{Name: "fuel pump modules"}
	db.Model(&Process{}).Create(&process12)

	process13 := Process{Name: "diesel fuel filters"}
	db.Model(&Process{}).Create(&process13)

	// Plant -----------------------------------------------------------------------------------------------------------------------------------------
	plant1 := Plant{Name: "Chongqing Chaoli Electric Appliance Co., Ltd."}
	db.Model(&Plant{}).Create(&plant1)

	plant2 := Plant{Name: "DENSO AIR SYSTEMS TIANJIN CO., LTD."}
	db.Model(&Plant{}).Create(&plant2)

	plant3 := Plant{Name: "DENSO (CHANGZHOU) FUEL INJECTION SYSTEM CO., LTD."}
	db.Model(&Plant{}).Create(&plant3)

	plant4 := Plant{Name: "DENSO FA SHANGHAI CO., LTD"}
	db.Model(&Plant{}).Create(&plant4)

	// Province -----------------------------------------------------------------------------------------------------------------------------------------
	province1 := Province{Name: "Chonburi"}
	db.Model(&Province{}).Create(&province1)

	province2 := Province{Name: "Samutprakarn"}
	db.Model(&Province{}).Create(&province2)

	// Country -----------------------------------------------------------------------------------------------------------------------------------------
	country1 := Country{Name: "China"}
	db.Model(&Country{}).Create(&country1)

	country2 := Country{Name: "South Korea"}
	db.Model(&Country{}).Create(&country2)

	country3 := Country{Name: "Taiwan"}
	db.Model(&Country{}).Create(&country3)

	country4 := Country{Name: "Japan"}
	db.Model(&Country{}).Create(&country4)

	country5 := Country{Name: "India"}
	db.Model(&Country{}).Create(&country5)
}
