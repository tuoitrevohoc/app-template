import { graphql, useLazyLoadQuery } from "react-relay";
import { HomeQuery } from "./__generated__/HomeQuery.graphql";

export function Home() {
  const data = useLazyLoadQuery<HomeQuery>(
    graphql`
      query HomeQuery {
        invoices {
          edges {
            node {
              id
            }
          }
        }
      }
    `,
    {}
  );

  return <div>Total invoices: {data.invoices.edges?.length}</div>;
}
