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

	// process14 := Process{Name: "air-conditioners and radiators"}
	// db.Model(&Plant{}).Create(&process14)

	// process15 := Process{Name: "thermal Products"}
	// db.Model(&Plant{}).Create(&process15)

	// process16 := Process{Name: "thermal Products"}
	// db.Model(&Plant{}).Create(&process16)

	// process17 := Process{Name: "thermal Products"}
	// db.Model(&Plant{}).Create(&process17)

	// process18 := Process{Name: "thermal Products"}
	// db.Model(&Plant{}).Create(&process18)

	// process19 := Process{Name: "electronic control products and components for automobiles"}
	// db.Model(&Plant{}).Create(&process19)

	// process20 := Process{Name: "HVAC (except for heat exchangers), condensers, electric fans and ECM (engine cooling module)"}
	// db.Model(&Plant{}).Create(&process20)

	// process21 := Process{Name: "automotive engine control system products and air-conditioning system products"}
	// db.Model(&Plant{}).Create(&process21)

	// process22 := Process{Name: "thermal Products"}
	// db.Model(&Plant{}).Create(&process22)

	// process23 := Process{Name: "car air-conditioners and compressors"}
	// db.Model(&Plant{}).Create(&process23)

	// process24 := Process{Name: "automotive components, non-automotive equipment and components"}
	// db.Model(&Plant{}).Create(&process24)

	// process25 := Process{Name: "wiper arms, wiper blades, and wiper linkages"}
	// db.Model(&Plant{}).Create(&process25)

	// process26 := Process{Name: "air-conditioners, radiators, bus air-conditioners, alternators, non-automotive products"}
	// db.Model(&Plant{}).Create(&process26)

	// process27 := Process{Name: "starters, alternators, Inverter, motor generator, electric power steering motor"}
	// db.Model(&Plant{}).Create(&process27)

	// process28 := Process{Name: "Car air conditioners, ratiators, electronic diesel/gasoline fuel injection components"}
	// db.Model(&Plant{}).Create(&process28)

	// process29 := Process{Name: "Instrument Cluster, Head-Up Display, Millimeter-Wave Radar Sensor, Vision Sensor ,Power Module,various Sensors components"}
	// db.Model(&Plant{}).Create(&process29)

	// process30 := Process{Name: "compact motors for wiper systems, power windows"}
	// db.Model(&Plant{}).Create(&process30)

	// process31 := Process{Name: "ignition components, Advanced safety products, valve train components, exhaust emission control components"}
	// db.Model(&Plant{}).Create(&process31)

	// process32 := Process{Name: "integrated circuits and electronic control components"}
	// db.Model(&Plant{}).Create(&process32)

	// process33 := Process{Name: "air conditioners and Water pump for Fuel-cell vehicle and Servo Motor module and natural refrigerant (CO2) heat-pump hot water supply system"}
	// db.Model(&Plant{}).Create(&process33)

	// process34 := Process{Name: "Inverters, ECUs and Power semiconductors"}
	// db.Model(&Plant{}).Create(&process34)

	// process35 := Process{Name: "machinery and tools"}
	// db.Model(&Plant{}).Create(&process35)

	// process36 := Process{Name: "compact motors for blowers, cooling fans"}
	// db.Model(&Plant{}).Create(&process36)

	// process37 := Process{Name: "electronic diesel/gasoline fuel injection systems"}
	// db.Model(&Plant{}).Create(&process37)

	// process38 := Process{Name: "automotive small motors, their systems and related components"}
	// db.Model(&Plant{}).Create(&process38)

	// process39 := Process{Name: "air-conditioners for buses, construction vehicles and agricultural vehicles, truck refrigerators, industrial air-conditioners, residential air-conditioners, and heat exchangers"}
	// db.Model(&Plant{}).Create(&process39)

	// process40 := Process{Name: "hoses, ducts, and pipes for car air-conditioning systems"}
	// db.Model(&Plant{}).Create(&process40)

	// process41 := Process{Name: "automotive fuel injection systems components, automotive air control system components"}
	// db.Model(&Plant{}).Create(&process41)

	// process42 := Process{Name: "automotive relays and electronic controllers"}
	// db.Model(&Plant{}).Create(&process42)

	// process43 := Process{Name: "automated production lines"}
	// db.Model(&Plant{}).Create(&process43)

	// process44 := Process{Name: "air-conditioning systems, engine cooling modules and fuel injection/supply systems for gasoline engines"}
	// db.Model(&Plant{}).Create(&process44)

	// process45 := Process{Name: "automotive semiconductor products"}
	// db.Model(&Plant{}).Create(&process45)

	// process46 := Process{Name: "automotive semiconductors and electronic device parts"}
	// db.Model(&Plant{}).Create(&process46)

	// process47 := Process{Name: "components for hybrid vehicles, starters, ignition systems, and intake air systems"}
	// db.Model(&Plant{}).Create(&process47)

	// process48 := Process{Name: "thermal systems products, diesel fuel injection systems components, intake air systems products, automotive wiper, and washer systems products"}
	// db.Model(&Plant{}).Create(&process48)

	// process49 := Process{Name: "automotive compact motors and their related components"}
	// db.Model(&Plant{}).Create(&process49)

	// process50 := Process{Name: "stamping parts for automobiles"}
	// db.Model(&Plant{}).Create(&process50)

	// process51 := Process{Name: "heat exchanger for trucks and agricultural machineries, air-conditioner for trucks"}
	// db.Model(&Plant{}).Create(&process51)

	// process52 := Process{Name: "Assembly of automotive components"}
	// db.Model(&Plant{}).Create(&process52)

	// process53 := Process{Name: "audio-visual products and automotive electronics"}
	// db.Model(&Plant{}).Create(&process53)

	// process54 := Process{Name: "resin parts for cars"}
	// db.Model(&Plant{}).Create(&process54)

	// process55 := Process{Name: "motorcycle components and automobile components"}
	// db.Model(&Plant{}).Create(&process55)

	// process56 := Process{Name: "automatic data capture equipment, industrial robots, programmable controllers and IoT, other apparatus and systems"}
	// db.Model(&Plant{}).Create(&process56)

	// process57 := Process{Name: "automotive wiper blades, wiper arms, wiper links"}
	// db.Model(&Plant{}).Create(&process57)

	// process58 := Process{Name: "automotive electric components"}
	// db.Model(&Plant{}).Create(&process58)

	// process59 := Process{Name: "automotive electric components"}
	// db.Model(&Plant{}).Create(&process59)

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

	plant48 := Plant{Name: "DENSO KATSUYAMA CO., LTD."}
	db.Model(&Plant{}).Create(&plant48)

	plant49 := Plant{Name: "DENSO KYUSHU CORPORATION"}
	db.Model(&Plant{}).Create(&plant49)

	plant50 := Plant{Name: "DENSO MIYAZAKI, INC."}
	db.Model(&Plant{}).Create(&plant50)

	plant51 := Plant{Name: "DENSO PRESS TECH CO., LTD."}
	db.Model(&Plant{}).Create(&plant51)

	plant52 := Plant{Name: "DENSO SANKYO CO., LTD."}
	db.Model(&Plant{}).Create(&plant52)

	plant53 := Plant{Name: "DENSO TAIYO CO., LTD."}
	db.Model(&Plant{}).Create(&plant53)

	plant54 := Plant{Name: "DENSO TEN Limited"}
	db.Model(&Plant{}).Create(&plant54)

	plant55 := Plant{Name: "DENSO TEN TECHNOSEPTA Limited"}
	db.Model(&Plant{}).Create(&plant55)

	plant56 := Plant{Name: "DENSO TRIM CORPORATION"}
	db.Model(&Plant{}).Create(&plant56)

	plant57 := Plant{Name: "DENSO WAVE INCORPORATED"}
	db.Model(&Plant{}).Create(&plant57)

	plant58 := Plant{Name: "DENSO WIPER SYSTEMS, INC."}
	db.Model(&Plant{}).Create(&plant58)

	plant59 := Plant{Name: "DENSO WISETECH CORPORATION"}
	db.Model(&Plant{}).Create(&plant59)

	plant60 := Plant{Name: "DENSO YAMAGATA CO., LTD."}
	db.Model(&Plant{}).Create(&plant60)

	plant61 := Plant{Name: "HAMADEN P.S CO.,LTD."}
	db.Model(&Plant{}).Create(&plant61)

	plant62 := Plant{Name: "HAMANAKODENSO CO., LTD."}
	db.Model(&Plant{}).Create(&plant62)

	plant63 := Plant{Name: "JECO CO., LTD."}
	db.Model(&Plant{}).Create(&plant63)

	plant64 := Plant{Name: "KYOSAN DENKI CO., LTD."}
	db.Model(&Plant{}).Create(&plant64)

	plant65 := Plant{Name: "KYOSAN TECS CO., LTD."}
	db.Model(&Plant{}).Create(&plant65)

	plant66 := Plant{Name: "MAULTECH CORPORATION"}
	db.Model(&Plant{}).Create(&plant66)

	plant67 := Plant{Name: "NAGANO JECO CO., LTD."}
	db.Model(&Plant{}).Create(&plant67)

	plant68 := Plant{Name: "SHIMIZU INDUSTRY CO., LTD."}
	db.Model(&Plant{}).Create(&plant68)

	plant69 := Plant{Name: "DENSO HARYANA PVT. LTD."}
	db.Model(&Plant{}).Create(&plant69)

	plant70 := Plant{Name: "DENSO INDIA PVT. LTD."}
	db.Model(&Plant{}).Create(&plant70)

	plant71 := Plant{Name: "DENSO KIRLOSKAR INDUSTRIES PVT. LTD."}
	db.Model(&Plant{}).Create(&plant71)

	plant72 := Plant{Name: "DENSO TEN MINDA INDIA Private Limited"}
	db.Model(&Plant{}).Create(&plant72)

	plant73 := Plant{Name: "DENSO THERMAL SYSTEMS PUNE PVT.LTD."}
	db.Model(&Plant{}).Create(&plant73)

	plant74 := Plant{Name: "PT. DENSO INDONESIA"}
	db.Model(&Plant{}).Create(&plant74)

	plant75 := Plant{Name: "PT. DENSO MANUFACTURING INDONESIA"}
	db.Model(&Plant{}).Create(&plant75)

	plant76 := Plant{Name: "PT. HAMADEN INDONESIA MANUFACTURING"}
	db.Model(&Plant{}).Create(&plant76)

	plant77 := Plant{Name: "AIR SYSTEMS (THAILAND) CO., LTD."}
	db.Model(&Plant{}).Create(&plant77)

	plant78 := Plant{Name: "DENSO ELECTRONICS (THAILAND) Co., LTD."}
	db.Model(&Plant{}).Create(&plant78)

	plant79 := Plant{Name: "DENSO (THAILAND) CO., LTD."}
	db.Model(&Plant{}).Create(&plant79)

	plant80 := Plant{Name: "DENSO TEN THAILAND Limited"}
	db.Model(&Plant{}).Create(&plant80)

	plant81 := Plant{Name: "DENSO INNOVATIVE MANUFACTURING SOLUTION ASIA CO., LTD."}
	db.Model(&Plant{}).Create(&plant81)

	plant82 := Plant{Name: "SIAM DENSO MANUFACTURING CO., LTD."}
	db.Model(&Plant{}).Create(&plant82)

	plant83 := Plant{Name: "IAM KYOSAN DENSO CO., LTD."}
	db.Model(&Plant{}).Create(&plant83)

	plant84 := Plant{Name: "DENSO MANUFACTURING VIETNAM CO., LTD."}
	db.Model(&Plant{}).Create(&plant84)

	plant85 := Plant{Name: "HAMADEN VIETNAM CO., LTD."}
	db.Model(&Plant{}).Create(&plant85)

	plant86 := Plant{Name: "DENSO (CAMBODIA) Co., Ltd."}
	db.Model(&Plant{}).Create(&plant86)

	plant87 := Plant{Name: "SANKYO RADIATOR (CAMBODIA) CO., LTD."}
	db.Model(&Plant{}).Create(&plant87)

	plant88 := Plant{Name: "DENSO (MALAYSIA) SDN. BHD. (198001003055 (56839-A))"}
	db.Model(&Plant{}).Create(&plant88)

	plant89 := Plant{Name: "DENSO WIPER SYSTEMS (MALAYSIA) SDN. BHD. (199501005004)"}
	db.Model(&Plant{}).Create(&plant89)

	plant90 := Plant{Name: "DENSO PHILIPPINES CORPORATION"}
	db.Model(&Plant{}).Create(&plant90)

	plant91 := Plant{Name: "Jeco Autoparts Philippines, Inc."}
	db.Model(&Plant{}).Create(&plant91)

	plant92 := Plant{Name: "DENSO TEN PHILIPPINES CORPORATION"}
	db.Model(&Plant{}).Create(&plant92)

	plant93 := Plant{Name: "DENSO THERMAL SYSTEMS MOROCCO S.A.R.L."}
	db.Model(&Plant{}).Create(&plant93)

	plant94 := Plant{Name: "DENSO AIR SYSTEMS KENTUCKY,INC."}
	db.Model(&Plant{}).Create(&plant94)

	plant95 := Plant{Name: "DENSO MANUFACTURING ARKANSAS, INC."}
	db.Model(&Plant{}).Create(&plant95)

	plant96 := Plant{Name: "DENSO MANUFACTURING ATHENS TENNESSEE, INC."}
	db.Model(&Plant{}).Create(&plant96)

	plant97 := Plant{Name: "DENSO MANUFACTURING MICHIGAN, INC."}
	db.Model(&Plant{}).Create(&plant97)

	plant98 := Plant{Name: "DENSO MANUFACTURING NORTH CAROLINA, INC."}
	db.Model(&Plant{}).Create(&plant98)

	plant99 := Plant{Name: "DENSO MANUFACTURING TENNESSEE, INC."}
	db.Model(&Plant{}).Create(&plant99)

	plant100 := Plant{Name: "DENSO TEN TECHNOSEPTA USA, Limited"}
	db.Model(&Plant{}).Create(&plant100)

	plant101 := Plant{Name: "YSTEX PRODUCTS ARKANSAS COMPANY"}
	db.Model(&Plant{}).Create(&plant101)

	plant102 := Plant{Name: "SYSTEX PRODUCTS CORPORATION"}
	db.Model(&Plant{}).Create(&plant102)

	plant103 := Plant{Name: "DENSO AIR SYSTEMS DE MEXICO S.A.DE C.V."}
	db.Model(&Plant{}).Create(&plant103)

	plant104 := Plant{Name: "DENSO MEXICO S.A. DE C.V."}
	db.Model(&Plant{}).Create(&plant104)

	plant105 := Plant{Name: "DENSO TEN DE MÉXICO, S.A. DE C.V."}
	db.Model(&Plant{}).Create(&plant105)

	plant106 := Plant{Name: "HAMADEN MEXICO S.A. DE C.V."}
	db.Model(&Plant{}).Create(&plant106)

	plant107 := Plant{Name: "DENSO MANUFACTURING CANADA, INC."}
	db.Model(&Plant{}).Create(&plant107)

	plant108 := Plant{Name: "DENSO DO BRASIL LTDA."}
	db.Model(&Plant{}).Create(&plant108)

	plant109 := Plant{Name: "DENSO INDUSTRIAL DA AMAZONIA LTDA."}
	db.Model(&Plant{}).Create(&plant109)

	plant110 := Plant{Name: "DENSO MAQUINAS ROTANTES do BRASIL LTDA."}
	db.Model(&Plant{}).Create(&plant110)

	plant111 := Plant{Name: "DENSO SISTEMAS TERMICOS do BRASIL LTDA."}
	db.Model(&Plant{}).Create(&plant111)

	plant112 := Plant{Name: "DENSO MANUFACTURING ARGENTINA S.A."}
	db.Model(&Plant{}).Create(&plant112)

	plant113 := Plant{Name: "DENSO MANUFACTURING HUNGARY LTD."}
	db.Model(&Plant{}).Create(&plant113)

	plant114 := Plant{Name: "DENSO CZECH s.r.o."}
	db.Model(&Plant{}).Create(&plant114)

	plant115 := Plant{Name: "DENSO MANUFACTURING CZECH s.r.o."}
	db.Model(&Plant{}).Create(&plant115)

	plant116 := Plant{Name: "LIPLASTEC s.r.o."}
	db.Model(&Plant{}).Create(&plant116)

	plant117 := Plant{Name: "DENSO POLAND Sp. z o.o."}
	db.Model(&Plant{}).Create(&plant117)

	plant118 := Plant{Name: "DENSO THERMAL SYSTEMS POLSKA Sp. z o.o."}
	db.Model(&Plant{}).Create(&plant118)

	plant119 := Plant{Name: "DENSO OTOMOTIV PARCALARI SANAYI A.S."}
	db.Model(&Plant{}).Create(&plant119)

	plant120 := Plant{Name: "DENSO MANUFACTURING UK LTD."}
	db.Model(&Plant{}).Create(&plant120)

	plant121 := Plant{Name: "DENSO MARSTON LTD."}
	db.Model(&Plant{}).Create(&plant121)

	plant122 := Plant{Name: "DENSO MANUFACTURING ITALIA S.p.A."}
	db.Model(&Plant{}).Create(&plant122)

	plant123 := Plant{Name: "DENSO THERMAL SYSTEMS S.p.A."}
	db.Model(&Plant{}).Create(&plant123)

	plant124 := Plant{Name: "DENSO BARCELONA S.A.U."}
	db.Model(&Plant{}).Create(&plant124)

	plant125 := Plant{Name: "DENSO SISTEMAS TERMICOS ESPANA S.A."}
	db.Model(&Plant{}).Create(&plant125)

	plant126 := Plant{Name: "DENSO TEN ESPAÑA, S.A."}
	db.Model(&Plant{}).Create(&plant126)

	plant127 := Plant{Name: "João de Deus & Filhos, S.A."}
	db.Model(&Plant{}).Create(&plant127)

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
