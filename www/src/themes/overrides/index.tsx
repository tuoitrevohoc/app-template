// third-party
import { Theme } from "@mui/material";
import { merge } from "lodash";

// project import
import CustomBadge from "./CustomBadge";
import CustomButton from "./CustomButton";
import CustomCardContent from "./CustomCardContent";
import Checkbox from "./CustomCheckbox";
import CustomChip from "./CustomChip";
import CustomIconButton from "./CustomIconButton";
import CustomInputLabel from "./CustomInputLabel";
import CustomLinearProgress from "./CustomLinearProgress";
import CustomLink from "./CustomLink";
import CustomListItemIcon from "./CustomListItemIcon";
import CustomOutlinedInput from "./CustomOutlinedInput";
import CustomTab from "./CustomTab";
import CustomTableCell from "./CustomTableCell";
import CustomTabs from "./CustomTabs";
import CustomTypography from "./CustomTypography";

// ==============================|| OVERRIDES - MAIN ||============================== //

export default function ComponentsOverrides(theme: Theme) {
  return merge(
    CustomButton(theme),
    CustomBadge(theme),
    CustomCardContent(),
    Checkbox(theme),
    CustomChip(theme),
    CustomIconButton(theme),
    CustomInputLabel(theme),
    CustomLinearProgress(),
    CustomLink(),
    CustomListItemIcon(),
    CustomOutlinedInput(theme),
    CustomTab(theme),
    CustomTableCell(theme),
    CustomTabs(),
    CustomTypography()
  );
}
