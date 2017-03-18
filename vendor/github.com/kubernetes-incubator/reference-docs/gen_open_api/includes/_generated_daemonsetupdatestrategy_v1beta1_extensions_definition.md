## DaemonSetUpdateStrategy v1beta1 extensions

Group        | Version     | Kind
------------ | ---------- | -----------
extensions | v1beta1 | DaemonSetUpdateStrategy





<aside class="notice">
Appears In  <a href="#daemonsetspec-v1beta1">DaemonSetSpec</a> </aside>

Field        | Description
------------ | -----------
rollingUpdate <br /> *[RollingUpdateDaemonSet](#rollingupdatedaemonset-v1beta1)*  | Rolling update config params. Present only if DaemonSetUpdateStrategy = RollingUpdate.
type <br /> *string*  | Type of daemon set update. Can be "RollingUpdate" or "OnDelete". Default is OnDelete.

