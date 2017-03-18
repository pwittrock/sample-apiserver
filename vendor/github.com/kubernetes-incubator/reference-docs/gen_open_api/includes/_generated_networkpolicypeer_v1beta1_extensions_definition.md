## NetworkPolicyPeer v1beta1 extensions

Group        | Version     | Kind
------------ | ---------- | -----------
extensions | v1beta1 | NetworkPolicyPeer





<aside class="notice">
Appears In  <a href="#networkpolicyingressrule-v1beta1">NetworkPolicyIngressRule</a> </aside>

Field        | Description
------------ | -----------
namespaceSelector <br /> *[LabelSelector](#labelselector-v1)*  | Selects Namespaces using cluster scoped-labels.  This matches all pods in all namespaces selected by this label selector. This field follows standard label selector semantics. If omitted, this selector selects no namespaces. If present but empty, this selector selects all namespaces.
podSelector <br /> *[LabelSelector](#labelselector-v1)*  | This is a label selector which selects Pods in this namespace. This field follows standard label selector semantics. If not provided, this selector selects no pods. If present but empty, this selector selects all pods in this namespace.

