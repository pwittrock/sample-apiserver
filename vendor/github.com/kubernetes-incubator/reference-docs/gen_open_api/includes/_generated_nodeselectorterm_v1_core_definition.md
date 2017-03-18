## NodeSelectorTerm v1 core

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | NodeSelectorTerm



A null or empty node selector term matches no objects.

<aside class="notice">
Appears In  <a href="#nodeselector-v1">NodeSelector</a>  <a href="#preferredschedulingterm-v1">PreferredSchedulingTerm</a> </aside>

Field        | Description
------------ | -----------
matchExpressions <br /> *[NodeSelectorRequirement](#nodeselectorrequirement-v1) array*  | Required. A list of node selector requirements. The requirements are ANDed.

