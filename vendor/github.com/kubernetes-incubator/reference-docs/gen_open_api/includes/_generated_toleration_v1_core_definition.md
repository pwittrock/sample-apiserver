## Toleration v1 core

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | Toleration



The pod this Toleration is attached to tolerates any taint that matches the triple <key,value,effect> using the matching operator <operator>.

<aside class="notice">
Appears In  <a href="#podspec-v1">PodSpec</a> </aside>

Field        | Description
------------ | -----------
effect <br /> *string*  | Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.
key <br /> *string*  | Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.
operator <br /> *string*  | Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.
tolerationSeconds <br /> *integer*  | TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.
value <br /> *string*  | Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.
