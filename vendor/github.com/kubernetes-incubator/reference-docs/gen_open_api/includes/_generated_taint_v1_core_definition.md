## Taint v1 core

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | Taint



The node this Taint is attached to has the effect "effect" on any pod that that does not tolerate the Taint.

<aside class="notice">
Appears In  <a href="#nodespec-v1">NodeSpec</a> </aside>

Field        | Description
------------ | -----------
effect <br /> *string*  | Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.
key <br /> *string*  | Required. The taint key to be applied to a node.
timeAdded <br /> *[Time](#time-v1)*  | TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.
value <br /> *string*  | Required. The taint value corresponding to the taint key.

