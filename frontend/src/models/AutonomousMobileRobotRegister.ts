import { CountriesInterface } from "./Country";
import { ProcessesInterface } from "./Process";
import { ProcessLinesInterface } from "./ProcessLine";
import { ProvincesInterface } from "./Province";

export interface AutonomousMobileRobotsInterface {
    ID: string,
    Name: string;
	SingleBoardComputerName: string;
	MicrocontrollerName:     string;
	MotorName:               string;
	MotorAmount:             number;
	EncoderName:             string;
	MotorDriverName:         string;
	BatteryName:             string;
	WheelName:               string;
	WheelAmount:             number;
	CameraName:              string;
	LidarName:               string;
	WiFiRouterName:          string;
	TemperatureSensorName:   string;

    CountryID: number;
    Country: CountriesInterface;

	ProvinceID: number;
	Province: ProvincesInterface;

    PlantID: number;
    Plant: CountriesInterface;

    ProcessID: number;
    Process: ProcessesInterface;

    ProcessLineID: number;
    ProcessLine: ProcessLinesInterface;
}
