/**
 * @generated SignedSource<<922c0749a9f53faaa8d650b36ff5b0ee>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Mutation } from 'relay-runtime';
export type DeleteInvoiceMutation$variables = {
  connections: ReadonlyArray<string>;
  id: string;
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
    "cacheID": "6b78eb63287d0163a6f9eca39ea84548",
    "id": null,
    "metadata": {},
    "name": "DeleteInvoiceMutation",
    "operationKind": "mutation",
    "text": "mutation DeleteInvoiceMutation(\n  $id: ID!\n) {\n  deleteInvoice(id: $id) {\n    id\n  }\n}\n"
  }
};
})();

(node as any).hash = "82b889025e0fcce25098a32646f94111";

export default node;
