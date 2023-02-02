import { DeleteOutline } from "@mui/icons-material";
import {
  Button,
  Card,
  CardActions,
  CardContent,
  Typography,
} from "@mui/material";
import { useCallback } from "react";
import { graphql, useFragment } from "react-relay";
import useDeleteInvoiceMutation from "./DeleteInvoiceMutation";
import { InvoiceCard_invoice$key } from "./__generated__/InvoiceCard_invoice.graphql";

interface Props {
  invoice: InvoiceCard_invoice$key;
  connectionKey: string;
}

export default function InvoiceCard(props: Props) {
  const invoice = useFragment(
    graphql`
      fragment InvoiceCard_invoice on Invoice {
        id
        invoicedTo
        title
        leetCodeLink
      }
    `,
    props.invoice
  );

  const [deleteInvoice] = useDeleteInvoiceMutation();

  const onDelete = useCallback(() => {
    deleteInvoice({
      variables: {
        id: invoice.id,
        connections: [props.connectionKey],
      },
    });
  }, [invoice.id, props.connectionKey]);

  return (
    <Card variant="outlined">
      <CardContent>
        <Typography sx={{ fontSize: 14 }} color="text.secondary" gutterBottom>
          To: {invoice.invoicedTo}
        </Typography>
        <Typography variant="h5" component="div">
          {invoice.title}
        </Typography>
        <Typography variant="body2">
          <a href={invoice.leetCodeLink}>Link</a>
        </Typography>
      </CardContent>
      <CardActions>
        <Button
          onClick={onDelete}
          color="warning"
          startIcon={<DeleteOutline />}
        >
          Delete
        </Button>
      </CardActions>
    </Card>
  );
}
