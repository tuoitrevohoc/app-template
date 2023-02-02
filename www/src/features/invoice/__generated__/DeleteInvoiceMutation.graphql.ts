/**
 * @generated SignedSource<<a64e05ad6439d8d50ad2c078b0960e12>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Mutation } from 'relay-runtime';
export type DeleteInvoiceMutation$variables = {
  connections: ReadonlyArray<string>;
  id: number;
};
export type DeleteInvoiceMutation$data = {
  readonly deleteInvoice: {
    readonly id: string;
  } | null;
};
export type DeleteInvoiceMutation = {
  response: DeleteInvoiceMutation$data;
  variables: DeleteInvoiceMutation$variables;
};

const node: ConcreteRequest = (function(){
var v0 = {
  "defaultValue": null,
  "kind": "LocalArgument",
  "name": "connections"
},
v1 = {
  "defaultValue": null,
  "kind": "LocalArgument",
  "name": "id"
},
v2 = [
  {
    "kind": "Variable",
    "name": "id",
    "variableName": "id"
  }
],
v3 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "id",
  "storageKey": null
};
return {
  "fragment": {
    "argumentDefinitions": [
      (v0/*: any*/),
      (v1/*: any*/)
    ],
    "kind": "Fragment",
    "metadata": null,
    "name": "DeleteInvoiceMutation",
    "selections": [
      {
        "alias": null,
        "args": (v2/*: any*/),
        "concreteType": "Invoice",
        "kind": "LinkedField",
        "name": "deleteInvoice",
        "plural": false,
        "selections": [
          (v3/*: any*/)
        ],
        "storageKey": null
      }
    ],
    "type": "Mutation",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [
      (v1/*: any*/),
      (v0/*: any*/)
    ],
    "kind": "Operation",
    "name": "DeleteInvoiceMutation",
    "selections": [
      {
        "alias": null,
        "args": (v2/*: any*/),
        "concreteType": "Invoice",
        "kind": "LinkedField",
        "name": "deleteInvoice",
        "plural": false,
        "selections": [
          (v3/*: any*/),
          {
            "alias": null,
            "args": null,
            "filters": null,
            "handle": "deleteEdge",
            "key": "",
            "kind": "ScalarHandle",
            "name": "id",
            "handleArgs": [
              {
                "kind": "Variable",
                "name": "connections",
                "variableName": "connections"
              }
            ]
          }
        ],
        "storageKey": null
      }
    ]
  },
  "params": {
    "cacheID": "c99ce096aa7983fb69199dcd4018e039",
    "id": null,
    "metadata": {},
    "name": "DeleteInvoiceMutation",
    "operationKind": "mutation",
    "text": "mutation DeleteInvoiceMutation(\n  $id: Int!\n) {\n  deleteInvoice(id: $id) {\n    id\n  }\n}\n"
  }
};
})();

(node as any).hash = "74a6bac509135cfcb083a280c77e698f";

export default node;
