import { graphql, useMutation } from "react-relay";
import { DeleteInvoiceMutation } from "./__generated__/DeleteInvoiceMutation.graphql";

export default function useDeleteInvoiceMutation() {
  return useMutation<DeleteInvoiceMutation>(graphql`
    mutation DeleteInvoiceMutation($id: ID!, $connections: [ID!]!) {
      deleteInvoice(id: $id) {
        id @deleteEdge(connections: $connections)
      }
    }
  `);
}
