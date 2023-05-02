import Typography from "@mui/material/Typography";
import { Paper } from "@mui/material";
function Login() {
    return (
        <div>
            <Paper elevation={3} sx={{ margin: 5, width: 550, height: 300 }}>
                <Typography textAlign="center">Login</Typography>
            </Paper>
        </div>
    );
}
export default Login;