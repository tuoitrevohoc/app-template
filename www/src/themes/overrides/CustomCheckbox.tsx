// ==============================|| OVERRIDES - CHECKBOX ||============================== //

import { Theme } from "@mui/material";

export default function CustomCheckbox(theme: Theme) {
  return {
    MuiCheckbox: {
      styleOverrides: {
        root: {
          color: theme.palette.secondary.light,
        },
      },
    },
  };
}
