// ==============================|| OVERRIDES - LINER PROGRESS ||============================== //

export default function CustomLinearProgress() {
  return {
    MuiLinearProgress: {
      styleOverrides: {
        root: {
          height: 6,
          borderRadius: 100,
        },
        bar: {
          borderRadius: 100,
        },
      },
    },
  };
}
