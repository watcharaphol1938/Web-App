import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';
import { UsersInterface } from "../models/IUser";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { Divider, Drawer, FormControl, FormHelperText, Grid, IconButton, InputAdornment, InputLabel, LinearProgress, LinearProgressProps, List, ListItem, ListItemButton, ListItemIcon, ListItemText, MenuItem, OutlinedInput, Paper, Select, TextField, styled, useTheme } from "@mui/material";
import { text } from "stream/consumers";
import { AutonomousMobileRobotsInterface } from "../models/AMRregister";
function AutonomousMobileRobot() {
    const [autonomousmobilerobots, setAutonomousMobileRobots] = React.useState<Partial<AutonomousMobileRobotsInterface>>({
        CountryID: 0, PlantID: 0, ProcessID: 0, ProcessLineID: 0,
        Name: "", SingleBoardComputerName: "", MicrocontrollerName: "", MotorName: "", MotorAmount: 0, EncoderName: "", MotorDriverName: "", BatteryName: "", WheelName: "", WheelAmount: 0, CameraName: "", LidarName: "", WiFiRouterName: "", TemperatureSensorName: "",
    });
    const [users, setUsers] = React.useState<UsersInterface[]>([]);
    // const getUsers = async () => {
    //     const apiUrl = "http://localhost:8080/users";
    //     const requestOptions = {
    //         method: "GET",
    //         headers: { "Content-Type": "application/json" },
    //     };
    //     fetch(apiUrl, requestOptions)
    //         .then((response) => response.json())
    //         .then((res) => {
    //             console.log(res.data);
    //             if (res.data) {
    //                 setUsers(res.data);
    //             }
    //         });
    // };
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
    useEffect(() => {
        getAutonomousMobileRobots();
    }, []);
    return (
        <div>
            <Paper elevation={3} sx={{ margin: 5, mx: 'auto', width: 928, height: 'auto' }}>
                <Typography textAlign="center">AMR</Typography>
                <Grid container spacing={1} sx={{ mx: 17, my: 2}}>
                    <Grid item xs={3}>
                        <p>Name</p>
                        <FormControl fullWidth>
                            <TextField id="Name"
                                value={autonomousmobilerobots.Name}
                                // onChange={handleInputChange}
                                label="Name"
                                variant="outlined"
                                size="small"
                                sx={{ width: '70%' }} />
                        </FormControl>
                    </Grid>
                    <Grid item xs={3}>
                        <p>Single-Board Computer</p>
                        <FormControl fullWidth>
                            <TextField id="Single-Board Computer"
                                value={autonomousmobilerobots.SingleBoardComputerName}
                                // onChange={handleInputChange}
                                label="Single-Board Computer"
                                variant="outlined"
                                size="small"
                                sx={{ width: '70%' }} />
                        </FormControl>
                    </Grid>
                    <Grid item xs={3}>
                        <p>Microcontroller</p>
                        <FormControl fullWidth>
                            <TextField id="Microcontroller"
                                value={autonomousmobilerobots.MicrocontrollerName}
                                // onChange={handleInputChange}
                                label="Microcontroller"
                                variant="outlined"
                                size="small"
                                sx={{ width: '70%' }} />
                        </FormControl>
                    </Grid>
                </Grid>
                <Grid container spacing={1} sx={{ mx: 17, my: 2}}>
                    <Grid item xs={3}>
                        <p>Motor</p>
                        <FormControl fullWidth>
                            <TextField id="Motor"
                                value={autonomousmobilerobots.MotorName}
                                // onChange={handleInputChange}
                                label="Motor"
                                variant="outlined"
                                size="small"
                                sx={{ width: '70%' }} />
                        </FormControl>
                    </Grid>
                    <Grid item xs={3}>
                        <p>Amount of Motor</p>
                        <FormControl fullWidth>
                            <TextField id="MotorAmount"
                                value={autonomousmobilerobots.MotorAmount}
                                // onChange={handleInputChange}
                                label="MotorAmount"
                                variant="outlined"
                                size="small"
                                sx={{ width: '70%' }} />
                        </FormControl>
                    </Grid>
                    <Grid item xs={3}>
                        <p>Encoder</p>
                        <FormControl fullWidth>
                            <TextField id="Encoder"
                                value={autonomousmobilerobots.EncoderName}
                                // onChange={handleInputChange}
                                label="Encoder"
                                variant="outlined"
                                size="small"
                                sx={{ width: '70%' }} />
                        </FormControl>
                    </Grid>
                </Grid>
                <Grid container spacing={1} sx={{ mx: 17, my: 2}}>
                    <Grid item xs={3}>
                        <p>Motor Driver</p>
                        <FormControl fullWidth>
                            <TextField id="Motor Driver"
                                value={autonomousmobilerobots.MotorDriverName}
                                // onChange={handleInputChange}
                                label="Motor Driver"
                                variant="outlined"
                                size="small"
                                sx={{ width: '70%' }} />
                        </FormControl>
                    </Grid>
                    <Grid item xs={3}>
                        <p>Wheel</p>
                        <FormControl fullWidth>
                            <TextField id="Wheel"
                                value={autonomousmobilerobots.WheelName}
                                // onChange={handleInputChange}
                                label="Wheel"
                                variant="outlined"
                                size="small"
                                sx={{ width: '70%' }} />
                        </FormControl>
                    </Grid>
                    <Grid item xs={3}>
                        <p>Amount of Wheel</p>
                        <FormControl fullWidth>
                            <TextField id="Amount of Wheel"
                                value={autonomousmobilerobots.WheelAmount}
                                // onChange={handleInputChange}
                                label="Amount of Wheel"
                                variant="outlined"
                                size="small"
                                sx={{ width: '70%' }} />
                        </FormControl>
                    </Grid>
                </Grid>
                <Grid container spacing={1} sx={{ mx: 17, my: 2}}>
                    <Grid item xs={3}>
                        <p>Battery</p>
                        <FormControl fullWidth>
                            <TextField id="Battery"
                                value={autonomousmobilerobots.BatteryName}
                                // onChange={handleInputChange}
                                label="Battery"
                                variant="outlined"
                                size="small"
                                sx={{ width: '70%' }} />
                        </FormControl>
                    </Grid>
                    <Grid item xs={3}>
                        <p>Camera</p>
                        <FormControl fullWidth>
                            <TextField id="Camera"
                                value={autonomousmobilerobots.CameraName}
                                // onChange={handleInputChange}
                                label="Camera"
                                variant="outlined"
                                size="small"
                                sx={{ width: '70%' }} />
                        </FormControl>
                    </Grid>
                    <Grid item xs={3}>
                        <p>LiDAR</p>
                        <FormControl fullWidth>
                            <TextField id="LiDAR"
                                value={autonomousmobilerobots.LidarName}
                                // onChange={handleInputChange}
                                label="LiDAR"
                                variant="outlined"
                                size="small"
                                sx={{ width: '70%' }} />
                        </FormControl>
                    </Grid>
                </Grid>
                <Grid container spacing={1} sx={{ mx: 17, my: 2}}>
                    <Grid item xs={3}>
                        <p>WiFi Router</p>
                        <FormControl fullWidth>
                            <TextField id="WiFi Router"
                                value={autonomousmobilerobots.WiFiRouterName}
                                // onChange={handleInputChange}
                                label="WiFi Router"
                                variant="outlined"
                                size="small"
                                sx={{ width: '70%' }} />
                        </FormControl>
                    </Grid>
                    <Grid item xs={3}>
                        <p>Temperature Sensor</p>
                        <FormControl fullWidth>
                            <TextField id="Temperature Sensor"
                                value={autonomousmobilerobots.TemperatureSensorName}
                                // onChange={handleInputChange}
                                label="Temperature Sensor"
                                variant="outlined"
                                size="small"
                                sx={{ width: '70%' }} />
                        </FormControl>
                    </Grid>
                    <Grid item xs={3}>
                        <p>Country</p>
                        <FormControl fullWidth>
                            <TextField id="Country"
                                value={autonomousmobilerobots.WiFiRouterName}
                                // onChange={handleInputChange}
                                label="Country"
                                variant="outlined"
                                size="small" 
                                sx={{ width: '70%' }}/>
                        </FormControl>
                    </Grid>
                </Grid>
                <Grid container spacing={1} sx={{ mx: 17, my: 2}}>
                    <Grid item xs={3}>
                        <p>Plant</p>
                        <FormControl fullWidth>
                            <TextField id="Plant"
                                value={autonomousmobilerobots.TemperatureSensorName}
                                // onChange={handleInputChange}
                                label="Plant"
                                variant="outlined"
                                size="small" 
                                sx={{ width: '70%' }}/>
                        </FormControl>
                    </Grid>
                    <Grid item xs={3}>
                        <p>Process</p>
                        <FormControl fullWidth>
                            <TextField id="Process"
                                value={autonomousmobilerobots.WiFiRouterName}
                                // onChange={handleInputChange}
                                label="Process"
                                variant="outlined"
                                size="small" 
                                sx={{ width: '70%' }}/>
                        </FormControl>
                    </Grid>
                    <Grid item xs={3}>
                        <p>Process Line</p>
                        <FormControl fullWidth>
                            <TextField id="Process Line"
                                value={autonomousmobilerobots.TemperatureSensorName}
                                // onChange={handleInputChange}
                                label="Process Line"
                                variant="outlined"
                                size="small" 
                                sx={{ width: '70%' }}/>
                        </FormControl>
                    </Grid>
                </Grid>
                <Button variant="outlined" sx={{ margin: 2 , mx: 18}}>Register</Button>
            </Paper>
        </div >
    );
}
export default AutonomousMobileRobot;
