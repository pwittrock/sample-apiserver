## NodeSelector v1 core

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | NodeSelector



A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.

<aside class="notice">
Appears In  <a href="#nodeaffinity-v1">NodeAffinity</a> </aside>

Field        | Description
------------ | -----------
nodeSelectorTerms <br /> *[NodeSelectorTerm](#nodeselectorterm-v1) array*  | Required. A list of node selector terms. The terms are ORed.

