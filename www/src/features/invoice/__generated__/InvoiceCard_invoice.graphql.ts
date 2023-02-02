/**
 * @generated SignedSource<<b96869d5dd595d7801ef86c878c1512d>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { Fragment, ReaderFragment } from 'relay-runtime';
import { FragmentRefs } from "relay-runtime";
export type InvoiceCard_invoice$data = {
  readonly id: string;
  readonly invoicedTo: string;
  readonly leetCodeLink: string;
  readonly title: string;
  readonly " $fragmentType": "InvoiceCard_invoice";
};
export type InvoiceCard_invoice$key = {
  readonly " $data"?: InvoiceCard_invoice$data;
  readonly " $fragmentSpreads": FragmentRefs<"InvoiceCard_invoice">;
};

const node: ReaderFragment = {
  "argumentDefinitions": [],
  "kind": "Fragment",
  "metadata": null,
  "name": "InvoiceCard_invoice",
  "selections": [
    {
      "alias": null,
      "args": null,
      "kind": "ScalarField",
      "name": "id",
      "storageKey": null
    },
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
  "type": "Invoice",
  "abstractKey": null
};

(node as any).hash = "412b5b18e2317cc950d506c666c82b1c";

export default node;
