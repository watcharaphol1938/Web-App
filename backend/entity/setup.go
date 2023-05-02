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
	database.AutoMigrate(&AutonomousMobileRobot{}, &ProcessLine{}, &Process{}, &Plant{}, &Country{})

	db = database

	// ProcessLine -----------------------------------------------------------------------------------------------------------------------------------------
	// Process -----------------------------------------------------------------------------------------------------------------------------------------
	// Plant -----------------------------------------------------------------------------------------------------------------------------------------
	plant1 := Plant{Name: "Chongqing Chaoli Electric Appliance Co., Ltd."}
	db.Model(&Plant{}).Create(&plant1)

	plant2 := Plant{Name: "DENSO AIR SYSTEMS TIANJIN CO., LTD."}
	db.Model(&Plant{}).Create(&plant2)

	plant3 := Plant{Name: "DENSO (CHANGZHOU) FUEL INJECTION SYSTEM CO., LTD."}
	db.Model(&Plant{}).Create(&plant3)

	plant4 := Plant{Name: "DENSO FA SHANGHAI CO., LTD"}
	db.Model(&Plant{}).Create(&plant4)

	plant5 := Plant{Name: "DENSO (GUANGZHOU NANSHA) CO., LTD."}
	db.Model(&Plant{}).Create(&plant5)

	plant6 := Plant{Name: "DENSO KOTEI AUTOMOTIVE ELECTRONICS (WUHAN) CO., LTD"}
	db.Model(&Plant{}).Create(&plant6)

	plant7 := Plant{Name: "DENSO MANUFACTURING (HANGZHOU) CO., LTD."}
	db.Model(&Plant{}).Create(&plant7)

	plant8 := Plant{Name: "DENSO MANUFACTURING (TIANJIN) MOTOR CO., LTD."}
	db.Model(&Plant{}).Create(&plant8)

	plant9 := Plant{Name: "DENSO TEN ELECTRONICS (WUXI) Limited"}
	db.Model(&Plant{}).Create(&plant9)

	plant10 := Plant{Name: "DENSO TEN PRECISION ELECTRONICS (TIANJIN) Limited"}
	db.Model(&Plant{}).Create(&plant10)

	plant11 := Plant{Name: "DENSO (TIANJIN) BODY PARTS CO., LTD."}
	db.Model(&Plant{}).Create(&plant11)

	plant12 := Plant{Name: "DENSO (TIANJIN) THERMAL PRODUCTS CO., LTD."}
	db.Model(&Plant{}).Create(&plant12)

	plant13 := Plant{Name: "GONGCHENG DENSO (CHONGQING) CO., LTD"}
	db.Model(&Plant{}).Create(&plant13)

	plant14 := Plant{Name: "GUANGZHOU DENSO CO., LTD."}
	db.Model(&Plant{}).Create(&plant14)

	plant15 := Plant{Name: "Liuzhou Chaoli Car Parts Manufacturing Co., Ltd."}
	db.Model(&Plant{}).Create(&plant15)

	plant16 := Plant{Name: "Nanjing Voaasco Auto Parts Co., Ltd."}
	db.Model(&Plant{}).Create(&plant16)

	plant17 := Plant{Name: "Ningbo Hangzhou Bay New District Chaoli Auto Parts Manufacturing Co., Ltd."}
	db.Model(&Plant{}).Create(&plant17)

	plant18 := Plant{Name: "Qingdao Chaoli Car Parts Manufacturing Co., Ltd."}
	db.Model(&Plant{}).Create(&plant18)

	plant19 := Plant{Name: "TIANJIN DENSO ELECTRONICS CO., LTD."}
	db.Model(&Plant{}).Create(&plant19)

	plant20 := Plant{Name: "TIANJIN DENSO ENGINE ELECTRICAL PRODUCTS CO., LTD."}
	db.Model(&Plant{}).Create(&plant20)

	plant21 := Plant{Name: "TIANJIN FAWER DENSO AIR-CONDITIONER CO., LTD."}
	db.Model(&Plant{}).Create(&plant21)

	plant22 := Plant{Name: "WUXI DENSO AUTOMOTIVE PRODUCTS CO., LTD."}
	db.Model(&Plant{}).Create(&plant22)

	plant23 := Plant{Name: "Yantai Chaoli Auto Electrical Parts Co., LTD"}
	db.Model(&Plant{}).Create(&plant23)

	plant24 := Plant{Name: "YANTAI SHOUGANG DENSO CO., LTD."}
	db.Model(&Plant{}).Create(&plant24)

	plant25 := Plant{Name: "DENSO KOREA CORPORATION"}
	db.Model(&Plant{}).Create(&plant25)

	plant26 := Plant{Name: "KOREA WIPER BLADE CO., LTD."}
	db.Model(&Plant{}).Create(&plant26)

	plant27 := Plant{Name: "DENSO TAIWAN CORP."}
	db.Model(&Plant{}).Create(&plant27)

	plant28 := Plant{Name: "Anjo Plant"}
	db.Model(&Plant{}).Create(&plant28)

	plant29 := Plant{Name: "Nishio Plant"}
	db.Model(&Plant{}).Create(&plant29)

	plant30 := Plant{Name: "Takatana Plant"}
	db.Model(&Plant{}).Create(&plant30)

	plant31 := Plant{Name: "Kosai Plant"}
	db.Model(&Plant{}).Create(&plant31)

	plant32 := Plant{Name: "Daian Plant"}
	db.Model(&Plant{}).Create(&plant32)

	plant33 := Plant{Name: "Kota Plant"}
	db.Model(&Plant{}).Create(&plant33)

	plant34 := Plant{Name: "Toyohashi Plant"}
	db.Model(&Plant{}).Create(&plant34)

	plant35 := Plant{Name: "Hirose Plant"}
	db.Model(&Plant{}).Create(&plant35)

	plant36 := Plant{Name: "Agui Plant"}
	db.Model(&Plant{}).Create(&plant36)

	plant37 := Plant{Name: "Toyohashi East Plant"}
	db.Model(&Plant{}).Create(&plant37)

	plant38 := Plant{Name: "Zenmyo Plant"}
	db.Model(&Plant{}).Create(&plant38)

	plant39 := Plant{Name: "APINES INC."}
	db.Model(&Plant{}).Create(&plant39)

	plant40 := Plant{Name: "DENSO AIRCOOL CORPORATION"}
	db.Model(&Plant{}).Create(&plant40)

	plant41 := Plant{Name: "DENSO AIR SYSTEMS CORPORATION"}
	db.Model(&Plant{}).Create(&plant41)

	plant42 := Plant{Name: "DENSO DAISHIN CORPORATION"}
	db.Model(&Plant{}).Create(&plant42)

	plant43 := Plant{Name: "DENSO ELECTRONICS CORPORATION"}
	db.Model(&Plant{}).Create(&plant43)

	plant44 := Plant{Name: "DENSO FA YAMAGATA CO., LTD."}
	db.Model(&Plant{}).Create(&plant44)

	plant45 := Plant{Name: "DENSO FUKUSHIMA CORPORATION"}
	db.Model(&Plant{}).Create(&plant45)

	plant46 := Plant{Name: "DENSO HOKKAIDO CORPORATION"}
	db.Model(&Plant{}).Create(&plant46)

	plant47 := Plant{Name: "DENSO IWATE CORPORATION"}
	db.Model(&Plant{}).Create(&plant47)

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

	country6 := Country{Name: "Indonesia"}
	db.Model(&Country{}).Create(&country6)

	country7 := Country{Name: "Thailand"}
	db.Model(&Country{}).Create(&country7)

	country8 := Country{Name: "Singapore"}
	db.Model(&Country{}).Create(&country8)

	country9 := Country{Name: "Vietnam"}
	db.Model(&Country{}).Create(&country9)

	country10 := Country{Name: "Cambodia"}
	db.Model(&Country{}).Create(&country10)

	country11 := Country{Name: "Malaysia"}
	db.Model(&Country{}).Create(&country11)

	country12 := Country{Name: "Philippines"}
	db.Model(&Country{}).Create(&country12)

	country13 := Country{Name: "United Arab Emirates"}
	db.Model(&Country{}).Create(&country13)

	country14 := Country{Name: "Australia"}
	db.Model(&Country{}).Create(&country14)

	country15 := Country{Name: "Morocco"}
	db.Model(&Country{}).Create(&country15)

	country16 := Country{Name: "South Africa"}
	db.Model(&Country{}).Create(&country16)

	country17 := Country{Name: "United State of America"}
	db.Model(&Country{}).Create(&country17)

	country18 := Country{Name: "Mexico"}
	db.Model(&Country{}).Create(&country18)

	country19 := Country{Name: "Canada"}
	db.Model(&Country{}).Create(&country19)

	country20 := Country{Name: "Brazil"}
	db.Model(&Country{}).Create(&country20)

	country21 := Country{Name: "Argentina"}
	db.Model(&Country{}).Create(&country21)

	country22 := Country{Name: "Hungary"}
	db.Model(&Country{}).Create(&country22)

	country23 := Country{Name: "Czech"}
	db.Model(&Country{}).Create(&country23)

	country24 := Country{Name: "Poland"}
	db.Model(&Country{}).Create(&country24)

	country25 := Country{Name: "Turkey"}
	db.Model(&Country{}).Create(&country25)

	country26 := Country{Name: "Russia"}
	db.Model(&Country{}).Create(&country26)

	country27 := Country{Name: "England"}
	db.Model(&Country{}).Create(&country27)

	country28 := Country{Name: "Sweden"}
	db.Model(&Country{}).Create(&country28)

	country29 := Country{Name: "Italy"}
	db.Model(&Country{}).Create(&country29)

	country30 := Country{Name: "Spain"}
	db.Model(&Country{}).Create(&country30)

	country31 := Country{Name: "Portugal"}
	db.Model(&Country{}).Create(&country31)

	country32 := Country{Name: "Netherlands"}
	db.Model(&Country{}).Create(&country32)

	country33 := Country{Name: "Germany"}
	db.Model(&Country{}).Create(&country33)

	country34 := Country{Name: "Belgium"}
	db.Model(&Country{}).Create(&country34)

	country35 := Country{Name: "France"}
	db.Model(&Country{}).Create(&country35)
}
