## HostPortRange v1beta1 extensions

Group        | Version     | Kind
------------ | ---------- | -----------
extensions | v1beta1 | HostPortRange



Host Port Range defines a range of host ports that will be enabled by a policy for pods to use.  It requires both the start and end to be defined.

<aside class="notice">
Appears In  <a href="#podsecuritypolicyspec-v1beta1">PodSecurityPolicySpec</a> </aside>

Field        | Description
------------ | -----------
max <br /> *integer*  | max is the end of the range, inclusive.
min <br /> *integer*  | min is the start of the range, inclusive.

