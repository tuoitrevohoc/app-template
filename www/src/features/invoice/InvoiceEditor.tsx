import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  TextField,
  Typography,
} from "@mui/material";
import { useCallback, useState } from "react";
import useCreateInvoiceMutation from "./CreateInvoiceMutation";

interface Props {
  isOpen: boolean;
  close(): void;
  connectionKey: string;
}

export default function InvoiceEditor(props: Props) {
  const [createInvoice, isCreating] = useCreateInvoiceMutation();
  const [title, setTitle] = useState("");
  const [leetCodeLink, setLeetCodeLink] = useState("");
  const [invoicedTo, setInvoicedTo] = useState("");

  const onSubmit = useCallback(() => {
    createInvoice({
      variables: {
        input: {
          title,
          leetCodeLink,
          invoicedTo,
        },
        connections: [props.connectionKey],
      },
      onCompleted: props.close,
    });
  }, [props.close, title, leetCodeLink, invoicedTo]);

  return (
    <Dialog open={props.isOpen} onClose={props.close}>
      <DialogTitle>Add New Invoice</DialogTitle>
      <DialogContent>
        <Typography variant="body2">
          Send this invoice to someone you hate
        </Typography>
        <TextField
          autoFocus
          margin="dense"
          id="title"
          label="Title"
          type="text"
          fullWidth
          variant="outlined"
          value={title}
          onChange={(event) => setTitle(event.target.value)}
        />
        <TextField
          autoFocus
          margin="dense"
          id="leetCodeLink"
          label="Leet Code Link"
          type="text"
          fullWidth
          variant="outlined"
          value={leetCodeLink}
          onChange={(event) => setLeetCodeLink(event.target.value)}
        />
        <TextField
          autoFocus
          margin="dense"
          id="invoiceTo"
          label="Send To"
          type="text"
          fullWidth
          variant="outlined"
          value={invoicedTo}
          onChange={(event) => setInvoicedTo(event.target.value)}
        />
      </DialogContent>
      <DialogActions>
        <Button onClick={props.close} color="secondary">
          Cancel
        </Button>
        <Button onClick={onSubmit}>Add</Button>
      </DialogActions>
    </Dialog>
  );
}
