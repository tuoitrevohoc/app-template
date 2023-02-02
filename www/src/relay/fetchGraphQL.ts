import { RequestParameters } from "relay-runtime";

export async function fetchGraphQL(
  request: RequestParameters,
  variables: Record<string, unknown> = {}
) {
  let endPoint = "/graphql/query";

  try {
    const response = await fetch(endPoint, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        query: request.text!,
        variables,
      }),
    });

    return await response.json();
  } catch (error) {
    console.log("Error fetching GraphQL", error);
  }
}
