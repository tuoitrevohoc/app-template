/**
 * @generated SignedSource<<ba5ba3e37c7c79a3691240358ce96729>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Mutation } from 'relay-runtime';
import { FragmentRefs } from "relay-runtime";
export type CreateInvoiceInput = {
  invoicedTo: string;
  leetCodeLink: string;
  title: string;
};
export type CreateInvoiceMutation$variables = {
  connections: ReadonlyArray<string>;
  input: CreateInvoiceInput;
};
export type CreateInvoiceMutation$data = {
  readonly createInvoice: {
    readonly id: string;
    readonly " $fragmentSpreads": FragmentRefs<"InvoiceCard_invoice">;
  } | null;
};
export type CreateInvoiceMutation = {
  response: CreateInvoiceMutation$data;
  variables: CreateInvoiceMutation$variables;
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
  "name": "input"
},
v2 = [
  {
    "kind": "Variable",
    "name": "input",
    "variableName": "input"
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
    "name": "CreateInvoiceMutation",
    "selections": [
      {
        "alias": null,
        "args": (v2/*: any*/),
        "concreteType": "Invoice",
        "kind": "LinkedField",
        "name": "createInvoice",
        "plural": false,
        "selections": [
          (v3/*: any*/),
          {
            "args": null,
            "kind": "FragmentSpread",
            "name": "InvoiceCard_invoice"
          }
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
    "name": "CreateInvoiceMutation",
    "selections": [
      {
        "alias": null,
        "args": (v2/*: any*/),
        "concreteType": "Invoice",
        "kind": "LinkedField",
        "name": "createInvoice",
        "plural": false,
        "selections": [
          (v3/*: any*/),
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "invoicedTo",
            "storageKey": null
          },
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "title",
            "storageKey": null
          },
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "leetCodeLink",
            "storageKey": null
          }
        ],
        "storageKey": null
      },
      {
        "alias": null,
        "args": (v2/*: any*/),
        "filters": null,
        "handle": "appendNode",
        "key": "",
        "kind": "LinkedHandle",
        "name": "createInvoice",
        "handleArgs": [
          {
            "kind": "Variable",
            "name": "connections",
            "variableName": "connections"
          },
          {
            "kind": "Literal",
            "name": "edgeTypeName",
            "value": "InvoiceEdge"
          }
        ]
      }
    ]
  },
  "params": {
    "cacheID": "50ee9b3b2d1aaf4aaf38484c3283038a",
    "id": null,
    "metadata": {},
    "name": "CreateInvoiceMutation",
    "operationKind": "mutation",
    "text": "mutation CreateInvoiceMutation(\n  $input: CreateInvoiceInput!\n) {\n  createInvoice(input: $input) {\n    id\n    ...InvoiceCard_invoice\n  }\n}\n\nfragment InvoiceCard_invoice on Invoice {\n  id\n  invoicedTo\n  title\n  leetCodeLink\n}\n"
  }
};
})();

(node as any).hash = "d2da980850398aafa8a9455d397feb50";

export default node;
