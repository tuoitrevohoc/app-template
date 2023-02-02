import { Box, Button, Container } from "@mui/material";
import { Stack } from "@mui/system";
import { useState } from "react";
import { graphql, useLazyLoadQuery } from "react-relay";
import InvoiceCard from "../../features/invoice/InvoiceCard";
import InvoiceEditor from "../../features/invoice/InvoiceEditor";
import { HomeQuery } from "./__generated__/HomeQuery.graphql";

export function Home() {
  const data = useLazyLoadQuery<HomeQuery>(
    graphql`
      query HomeQuery($first: Int!) {
        invoices(first: $first) @connection(key: "Home__invoices") {
          __id
          edges {
            node {
              id
              ...InvoiceCard_invoice
            }
          }
        }
      }
    `,
    { first: 100 }
  );

  const [isDialogOpen, setIsDialogOpen] = useState(false);

  return (
    <Stack sx={{ maxWidth: 450, margin: "auto auto" }}>
      <h1>Total invoices: {data.invoices.edges?.length}</h1>
      <Stack sx={{ gap: 1 }}>
        {data.invoices.edges?.map(
          (edge) =>
            edge?.node && (
              <InvoiceCard
                key={edge.node.id}
                invoice={edge.node}
                connectionKey={data.invoices.__id}
              />
            )
        )}
      </Stack>
      <Box>
        <Button variant="contained" onClick={() => setIsDialogOpen(true)}>
          Add Invoice
        </Button>
      </Box>
      <InvoiceEditor
        isOpen={isDialogOpen}
        close={() => setIsDialogOpen(false)}
        connectionKey={data.invoices.__id}
      />
    </Stack>
  );
}
