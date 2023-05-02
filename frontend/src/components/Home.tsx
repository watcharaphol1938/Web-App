import Typography from "@mui/material/Typography";
import { Link as RouterLink } from "react-router-dom";
import { IconButton, Paper } from "@mui/material";
function Home() {
    return (
        <div>
            <Paper elevation={3} sx={{ margin: 5, width: 550, height: 300 }}>
                <Typography textAlign="center">Home</Typography>
                <IconButton size="small" component={RouterLink} to="/dashboard">Denso</IconButton>
            </Paper>
        </div>
    );
}
export default Home;