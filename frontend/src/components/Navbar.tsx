import React, { useState } from "react";
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import MenuItem from "@mui/material/MenuItem";
import { ThemeProvider, createTheme } from "@mui/material/styles";
import { red } from "@mui/material/colors";
import Drawer from "@mui/material/Drawer";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemText from "@mui/material/ListItemText";
import { Button, Divider, ListItemButton } from "@mui/material";
function Navbar() {
    const [openDrawer, setOpenDrawer] = useState(false);
    const menu = [
        { name: "Country", path: "/country" },
        { name: "Plant", path: "/plant" },
        { name: "Process", path: "/process" },
        { name: "AMR", path: "/amr" },
    ]
    const topic = [
        { name: 'Dashboard', path: "/dashboard" },
        { name: 'Analytics', path: "/analytics" },
        { name: 'Setting', path: "/setting" },
        { name: 'About', path: "/about" },
    ];
    const theme = createTheme({
        palette: {
            primary: {
                main: red[500],
            },
            secondary: {
                main: '#ff1744',
            },
        },
    });
    const handleDrawerToggle = () => {
        setOpenDrawer((prevState) => !prevState);
    };
    // const drawerWidth = 240;
    const drawer = (
        <Box onClick={handleDrawerToggle} sx={{ textAlign: 'center' }}>
            <Typography variant="h6" sx={{ my: 2 }}>
                Denso
            </Typography>
            <Divider />
            <List>
                {menu.map((menu, index) => (
                    <ListItem key={index} >
                        <ListItemButton sx={{ textAlign: 'center', paddingX: 5 }} component={RouterLink} to={menu.path}>
                            <ListItem>{menu.name}</ListItem>
                        </ListItemButton>
                    </ListItem>
                ))}
            </List>
        </Box>
    );
    return (
        <Box sx={{ flexGrow: 1 }}>
            <ThemeProvider theme={theme}>
                <AppBar position="static">
                    <Toolbar>
                        <IconButton size="medium" sx={{ color: "inherit" }} component={RouterLink} to="/">DENSO</IconButton>
                        <IconButton
                            size="large"
                            edge="start"
                            color="inherit"
                            onClick={handleDrawerToggle}
                            aria-label="menu"
                            sx={{ ml: 1 }}
                        >
                            <MenuIcon />
                        </IconButton>
                        <List>
                            {topic.map((topic, index) => (
                                <List key={index} component={RouterLink} to={topic.path}>
                                    <IconButton size="small" sx={{ color: "white" }}>{topic.name}</IconButton>
                                </List>
                            ))}
                        </List>
                        {/* <Drawer open={openDrawer} onClose={toggleDrawer(false)}>
                            <List
                                onClick={toggleDrawer(false)}
                                onKeyDown={toggleDrawer(false)}
                            >
                                {menu.map((item, index) => (
                                    <ListItem key={index} button
                                        component={RouterLink}
                                        to={item.path}>
                                        <ListItemText>{item.name}</ListItemText>
                                    </ListItem>)
                                )}
                            </List>
                        </Drawer> */}
                    </Toolbar>
                </AppBar>
                <Box component="nav">
                    <Drawer
                        // container={container}
                        variant="temporary"
                        open={openDrawer}
                        onClose={handleDrawerToggle}
                        ModalProps={{
                            keepMounted: true, // Better open performance on mobile.
                        }}
                    // sx={{
                    //     display: { xs: 'block', sm: 'none' },
                    //     '& .MuiDrawer-paper': { boxSizing: 'border-box', width: drawerWidth },
                    // }}
                    >
                        {drawer}
                    </Drawer>
                </Box>
            </ThemeProvider>
        </Box>
    );
}
export default Navbar;
