## PreferredSchedulingTerm v1 core

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | PreferredSchedulingTerm



An empty preferred scheduling term matches all objects with implicit weight 0 (i.e. it's a no-op). A null preferred scheduling term matches no objects (i.e. is also a no-op).

<aside class="notice">
Appears In  <a href="#nodeaffinity-v1">NodeAffinity</a> </aside>

Field        | Description
------------ | -----------
preference <br /> *[NodeSelectorTerm](#nodeselectorterm-v1)*  | A node selector term, associated with the corresponding weight.
weight <br /> *integer*  | Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.

