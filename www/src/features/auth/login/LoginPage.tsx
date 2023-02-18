import {
  Button,
  Card,
  CardContent,
  Checkbox,
  InputLabel,
  OutlinedInput,
  Stack,
  Typography,
} from "@mui/material";
import { Link } from "found";

export default function LoginPage() {
  return (
    <Stack
      sx={{ width: "100%", height: "100vh" }}
      alignItems="center"
      justifyContent="center"
    >
      <Card sx={{ width: 500 }}>
        <CardContent>
          <Stack gap={3}>
            <Stack direction="row" justifyContent="space-between">
              <Typography variant="h4">Login</Typography>
              <Link to="/register">Don't have an account?</Link>
            </Stack>

            <Stack gap={2}>
              <Stack gap={1}>
                <InputLabel>Email Address:</InputLabel>
                <OutlinedInput placeholder="Enter email address" />
              </Stack>
              <Stack gap={1}>
                <InputLabel>Password:</InputLabel>
                <OutlinedInput type="password" placeholder="Enter password" />
              </Stack>
            </Stack>
            <Stack direction="row" alignItems="center">
              <Checkbox />
              <InputLabel sx={{ flexGrow: 1 }}>Keep me sign in</InputLabel>
              <Link to="/forgot-password">Forgot password</Link>
            </Stack>
            <Stack>
              <Button variant="contained" color="primary" fullWidth>
                Log in
              </Button>
            </Stack>
          </Stack>
        </CardContent>
      </Card>
    </Stack>
  );
}
