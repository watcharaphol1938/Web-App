import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import { FormControl, Grid, Paper, Select, SelectChangeEvent, TextField } from "@mui/material";
import { AutonomousMobileRobotsInterface } from "../models/AutonomousMobileRobotRegister";
import { CountriesInterface } from "../models/Country";
import AutonomousMobileRobotRegisters from "./AutonomousMobileRobotRegister";
import { ProvincesInterface } from "../models/Province";
import { PlantsInterface } from "../models/Plant";
import { ProcessesInterface } from "../models/Process";
import { ProcessLinesInterface } from "../models/ProcessLine";
function AutonomousMobileRobotCreate() {
    const [autonomousmobilerobots, setAutonomousMobileRobots] = React.useState<Partial<AutonomousMobileRobotsInterface>>({
        CountryID: 0, PlantID: 0, ProcessID: 0, ProcessLineID: 0, ProvinceID: 0,
        Name: "", SingleBoardComputerName: "", MicrocontrollerName: "", MotorName: "", MotorAmount: 0, EncoderName: "", MotorDriverName: "", BatteryName: "", WheelName: "", WheelAmount: 0, CameraName: "", LidarName: "", WiFiRouterName: "", TemperatureSensorName: "",
    });
    const [countries, setCountries] = React.useState<CountriesInterface[]>([])
    const [provinces, setProvinces] = React.useState<CountriesInterface[]>([])
    const [plants, setPlants] = React.useState<CountriesInterface[]>([])
    const [processes, setProcesses] = React.useState<CountriesInterface[]>([])
    const [processlines, setProcessLines] = React.useState<CountriesInterface[]>([])

    const getAutonomousMobileRobots = async () => {
        const apiUrl = "http://localhost:8080/amrs";
        const requestOptions = {
            method: "GET",
            headers: { "Content-Type": "application/json" },
        };
        fetch(apiUrl, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                console.log(res.data);
                if (res.data) {
                    setAutonomousMobileRobots(res.data);
                }
            });
    };
    const getCountry = async () => {
        const apiUrl = "http://localhost:8080/countries";
        const requestOptions = {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                Authorization: `Bearer ${localStorage.getItem("token")}`
            }
        };

        // หลังบ้านรับ request มา
        // หลังบ้าน check data
        fetch(apiUrl, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setCountries(res.data);  // ข้อมูลถูกต้อง setPrefixes จะ set ค่าไปที่ตัวแปร prefixes
                } else {
                    console.log(res.error); // ข้อมูลไม่ถูกต้อง จะแสดงค่า error ที่ console
                }
            });
    }
    const getProvince = async () => {
        const apiUrl = "http://localhost:8080/provinces";
        const requestOptions = {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                Authorization: `Bearer ${localStorage.getItem("token")}`
            }
        };

        // หลังบ้านรับ request มา
        // หลังบ้าน check data
        fetch(apiUrl, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setProvinces(res.data);  // ข้อมูลถูกต้อง setPrefixes จะ set ค่าไปที่ตัวแปร prefixes
                } else {
                    console.log(res.error); // ข้อมูลไม่ถูกต้อง จะแสดงค่า error ที่ console
                }
            });
    }
    const getPlant = async () => {
        const apiUrl = "http://localhost:8080/plants";
        const requestOptions = {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                Authorization: `Bearer ${localStorage.getItem("token")}`
            }
        };

        // หลังบ้านรับ request มา
        // หลังบ้าน check data
        fetch(apiUrl, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setPlants(res.data);  // ข้อมูลถูกต้อง setPrefixes จะ set ค่าไปที่ตัวแปร prefixes
                } else {
                    console.log(res.error); // ข้อมูลไม่ถูกต้อง จะแสดงค่า error ที่ console
                }
            });
    }
    const getProcess = async () => {
        const apiUrl = "http://localhost:8080/processes";
        const requestOptions = {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                Authorization: `Bearer ${localStorage.getItem("token")}`
            }
        };

        // หลังบ้านรับ request มา
        // หลังบ้าน check data
        fetch(apiUrl, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setProcesses(res.data);  // ข้อมูลถูกต้อง setPrefixes จะ set ค่าไปที่ตัวแปร prefixes
                } else {
                    console.log(res.error); // ข้อมูลไม่ถูกต้อง จะแสดงค่า error ที่ console
                }
            });
    }
    const getProcessLine = async () => {
        const apiUrl = "http://localhost:8080/processlines";
        const requestOptions = {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                Authorization: `Bearer ${localStorage.getItem("token")}`
            }
        };

        // หลังบ้านรับ request มา
        // หลังบ้าน check data
        fetch(apiUrl, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setProcessLines(res.data);  // ข้อมูลถูกต้อง setPrefixes จะ set ค่าไปที่ตัวแปร prefixes
                } else {
                    console.log(res.error); // ข้อมูลไม่ถูกต้อง จะแสดงค่า error ที่ console
                }
            });
    }
    const handleSelectChange = (event: SelectChangeEvent<number>) => {
        const name = event.target.name as keyof typeof AutonomousMobileRobotRegisters;
        const { value } = event.target;
        setAutonomousMobileRobots({ ...autonomousmobilerobots, [name]: value });
    }
    const handleInputChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
        const id = event.target.id as keyof typeof AutonomousMobileRobotCreate;
        const { value } = event.target;
        setAutonomousMobileRobots({ ...autonomousmobilerobots, [id]: value });
    };
    useEffect(() => {
        getAutonomousMobileRobots();
        getCountry();
        getProvince();
        getPlant();
        getProcess();
        getProcessLine();
    }, []);
    return (
        <div>
            <Paper elevation={3} sx={{ margin: 5, mx: 'auto', width: 928, height: 'auto' }}>
                <Typography textAlign="center">AMR</Typography>
                <Grid container spacing={1} sx={{ mx: 17, my: 2 }}>
                    <Grid item xs={4}>
                        <p>Name</p>
                        <FormControl fullWidth>
                            <TextField id="Name"
                                value={autonomousmobilerobots.Name}
                                // onChange={handleInputChange}
                                label="Name"
                                variant="outlined"
                                size="small"
                            // sx={{ width: '70%' }} 
                            />
                        </FormControl>
                    </Grid>
                    <Grid item xs={4}>
                        <p>Single-Board Computer</p>
                        <FormControl fullWidth>
                            <TextField id="Single-Board Computer"
                                value={autonomousmobilerobots.SingleBoardComputerName}
                                // onChange={handleInputChange}
                                label="Single-Board Computer"
                                variant="outlined"
                                size="small"
                            // sx={{ width: '70%' }} 
                            />
                        </FormControl>
                    </Grid>
                </Grid>
                <Grid container spacing={1} sx={{ mx: 17, my: 2 }}>
                    <Grid item xs={4}>
                        <p>Microcontroller</p>
                        <FormControl fullWidth>
                            <TextField id="Microcontroller"
                                value={autonomousmobilerobots.MicrocontrollerName}
                                // onChange={handleInputChange}
                                label="Microcontroller"
                                variant="outlined"
                                size="small"
                            // sx={{ width: '70%' }} 
                            />
                        </FormControl>
                    </Grid>
                    <Grid item xs={4}>
                        <p>Encoder</p>
                        <FormControl fullWidth>
                            <TextField id="Encoder"
                                value={autonomousmobilerobots.EncoderName}
                                // onChange={handleInputChange}
                                label="Encoder"
                                variant="outlined"
                                size="small"
                            // sx={{ width: '70%' }}
                            />
                        </FormControl>
                    </Grid>
                </Grid>
                <Grid container spacing={1} sx={{ mx: 17, my: 2 }}>
                    <Grid item xs={4}>
                        <p>Motor</p>
                        <FormControl fullWidth>
                            <TextField id="Motor"
                                value={autonomousmobilerobots.MotorName}
                                // onChange={handleInputChange}
                                label="Motor"
                                variant="outlined"
                                size="small"
                            // sx={{ width: '70%' }} 
                            />
                        </FormControl>
                    </Grid>
                    <Grid item xs={4}>
                        <p>Amount of Motor</p>
                        <FormControl fullWidth>
                            <TextField id="MotorAmount"
                                value={autonomousmobilerobots.MotorAmount}
                                // onChange={handleInputChange}
                                label="MotorAmount"
                                variant="outlined"
                                size="small"
                            // sx={{ width: '70%' }} 
                            />
                        </FormControl>
                    </Grid>
                </Grid>
                <Grid container spacing={1} sx={{ mx: 17, my: 2 }}>
                    <Grid item xs={4}>
                        <p>Motor Driver</p>
                        <FormControl fullWidth>
                            <TextField id="Motor Driver"
                                value={autonomousmobilerobots.MotorDriverName}
                                // onChange={handleInputChange}
                                label="Motor Driver"
                                variant="outlined"
                                size="small"
                            // sx={{ width: '70%' }} 
                            />
                        </FormControl>
                    </Grid>
                    <Grid item xs={4}>
                        <p>Battery</p>
                        <FormControl fullWidth>
                            <TextField id="Battery"
                                value={autonomousmobilerobots.BatteryName}
                                // onChange={handleInputChange}
                                label="Battery"
                                variant="outlined"
                                size="small"
                            // sx={{ width: '70%' }} 
                            />
                        </FormControl>
                    </Grid>
                </Grid>
                <Grid container spacing={1} sx={{ mx: 17, my: 2 }}>
                    <Grid item xs={4}>
                        <p>Wheel</p>
                        <FormControl fullWidth>
                            <TextField id="Wheel"
                                value={autonomousmobilerobots.WheelName}
                                // onChange={handleInputChange}
                                label="Wheel"
                                variant="outlined"
                                size="small"
                            // sx={{ width: '70%' }} 
                            />
                        </FormControl>
                    </Grid>
                    <Grid item xs={4}>
                        <p>Amount of Wheel</p>
                        <FormControl fullWidth>
                            <TextField id="Amount of Wheel"
                                value={autonomousmobilerobots.WheelAmount}
                                // onChange={handleInputChange}
                                label="Amount of Wheel"
                                variant="outlined"
                                size="small"
                            // sx={{ width: '70%' }} 
                            />
                        </FormControl>
                    </Grid>
                </Grid>
                <Grid container spacing={1} sx={{ mx: 17, my: 2 }}>
                    <Grid item xs={4}>
                        <p>Camera</p>
                        <FormControl fullWidth>
                            <TextField id="Camera"
                                value={autonomousmobilerobots.CameraName}
                                // onChange={handleInputChange}
                                label="Camera"
                                variant="outlined"
                                size="small"
                            // sx={{ width: '70%' }} 
                            />
                        </FormControl>
                    </Grid>
                    <Grid item xs={4}>
                        <p>LiDAR</p>
                        <FormControl fullWidth>
                            <TextField id="LiDAR"
                                value={autonomousmobilerobots.LidarName}
                                // onChange={handleInputChange}
                                label="LiDAR"
                                variant="outlined"
                                size="small"
                            // sx={{ width: '70%' }} 
                            />
                        </FormControl>
                    </Grid>
                </Grid>
                <Grid container spacing={1} sx={{ mx: 17, my: 2 }}>
                    <Grid item xs={4}>
                        <p>WiFi Router</p>
                        <FormControl fullWidth>
                            <TextField id="WiFi Router"
                                value={autonomousmobilerobots.WiFiRouterName}
                                // onChange={handleInputChange}
                                label="WiFi Router"
                                variant="outlined"
                                size="small"
                            // sx={{ width: '70%' }} 
                            />
                        </FormControl>
                    </Grid>
                    <Grid item xs={4}>
                        <p>Temperature Sensor</p>
                        <FormControl fullWidth>
                            <TextField id="Temperature Sensor"
                                value={autonomousmobilerobots.TemperatureSensorName}
                                // onChange={handleInputChange}
                                label="Temperature Sensor"
                                variant="outlined"
                                size="small"
                            // sx={{ width: '70%' }} 
                            />
                        </FormControl>
                    </Grid>
                </Grid>
                <Grid container spacing={1} sx={{ mx: 17, my: 2 }}>
                    <Grid item xs={4}>
                        <p>Country</p>
                        <FormControl fullWidth>
                            {/* <InputLabel>Country</InputLabel> */}
                            <Select
                                id="CountryID"
                                title="CountryID"
                                inputProps={{
                                    name: "CountryID"
                                }}
                                value={autonomousmobilerobots.CountryID || 0}
                                onChange={handleSelectChange}
                                native
                                autoFocus
                                size="small"
                            >
                                <option key={0} value={0}></option>
                                {countries.map((item: CountriesInterface) => (
                                    <option key={item.ID} value={item.ID}>{item.Name}</option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={4}>
                        <p>Province</p>
                        <FormControl fullWidth>
                            {/* <InputLabel>Province</InputLabel> */}
                            <Select
                                id="ProvinceID"
                                title="ProvinceID"
                                inputProps={{
                                    name: "ProvinceID"
                                }}
                                value={autonomousmobilerobots.ProvinceID || 0}
                                onChange={handleSelectChange}
                                native
                                autoFocus
                                size="small"
                            >
                                <option key={0} value={0}></option>
                                {provinces.map((item: ProvincesInterface) => (
                                    <option key={item.ID} value={item.ID}>{item.Name}</option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                </Grid>
                <Grid container spacing={1} sx={{ mx: 13, my: 2 }}>
                    <Grid item xs={3}>
                        <p>Plant</p>
                        <FormControl fullWidth>
                            {/* <InputLabel>Plant</InputLabel> */}
                            <Select
                                id="PlantID"
                                title="PlantID"
                                inputProps={{
                                    name: "PlantID"
                                }}
                                value={autonomousmobilerobots.PlantID || 0}
                                onChange={handleSelectChange}
                                native
                                autoFocus
                                size="small"
                            >
                                <option key={0} value={0}></option>
                                {plants.map((item: PlantsInterface) => (
                                    <option key={item.ID} value={item.ID}>{item.Name}</option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={3}>
                        <p>Process</p>
                        <FormControl fullWidth>
                            {/* <InputLabel>Process</InputLabel> */}
                            <Select
                                id="ProcessID"
                                title="ProcessID"
                                inputProps={{
                                    name: "ProcessID"
                                }}
                                value={autonomousmobilerobots.ProcessID || 0}
                                onChange={handleSelectChange}
                                native
                                autoFocus
                                size="small"
                            >
                                <option key={0} value={0}></option>
                                {processes.map((item: ProcessesInterface) => (
                                    <option key={item.ID} value={item.ID}>{item.Name}</option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={3}>
                        <p>Process Line</p>
                        <FormControl fullWidth>
                            {/* <InputLabel>Process</InputLabel> */}
                            <Select
                                id="ProcessLineID"
                                title="ProcessLineID"
                                inputProps={{
                                    name: "ProcessLineID"
                                }}
                                value={autonomousmobilerobots.ProcessLineID || 0}
                                onChange={handleSelectChange}
                                native
                                autoFocus
                                size="small"
                            >
                                <option key={0} value={0}></option>
                                {processlines.map((item: ProcessLinesInterface) => (
                                    <option key={item.ID} value={item.ID}>{item.Name}</option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                </Grid>
                {/* <Grid container spacing={1} sx={{ mx: 17, my: 2 }}>

                </Grid> */}
                <Button variant="outlined" sx={{ margin: 2, mx: 18 }}>Register</Button>
            </Paper>
        </div >
    );
}
export default AutonomousMobileRobotCreate;
