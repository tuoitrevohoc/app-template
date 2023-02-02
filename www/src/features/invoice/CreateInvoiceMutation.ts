import { graphql, useMutation } from "react-relay";
import { CreateInvoiceMutation } from "./__generated__/CreateInvoiceMutation.graphql";

export default function useCreateInvoiceMutation() {
  return useMutation<CreateInvoiceMutation>(
    graphql`
      mutation CreateInvoiceMutation(
        $input: CreateInvoiceInput!
        $connections: [ID!]!
      ) {
        createInvoice(input: $input)
          @appendNode(connections: $connections, edgeTypeName: "InvoiceEdge") {
          id
          ...InvoiceCard_invoice
        }
      }
    `
  );
}
