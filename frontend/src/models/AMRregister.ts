import { CountriesInterface } from "./Country";
import { ProcessesInterface } from "./Process";
import { ProcessLinesInterface } from "./ProcessLine";

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

    PlantID: number;
    Plant: CountriesInterface;

    ProcessID: number;
    Process: ProcessesInterface;

    ProcessLineID: number;
    ProcessLine: ProcessLinesInterface;
}
