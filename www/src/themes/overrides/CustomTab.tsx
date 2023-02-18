// ==============================|| OVERRIDES - TAB ||============================== //

import { Theme } from "@mui/material";

export default function CustomTab(theme: Theme) {
  return {
    MuiTab: {
      styleOverrides: {
        root: {
          minHeight: 46,
          color: theme.palette.text.primary,
        },
      },
    },
  };
}
