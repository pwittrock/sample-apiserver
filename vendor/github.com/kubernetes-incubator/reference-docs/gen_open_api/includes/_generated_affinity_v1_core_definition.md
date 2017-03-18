## Affinity v1 core

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | Affinity



Affinity is a group of affinity scheduling rules.

<aside class="notice">
Appears In  <a href="#podspec-v1">PodSpec</a> </aside>

Field        | Description
------------ | -----------
nodeAffinity <br /> *[NodeAffinity](#nodeaffinity-v1)*  | Describes node affinity scheduling rules for the pod.
podAffinity <br /> *[PodAffinity](#podaffinity-v1)*  | Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).
podAntiAffinity <br /> *[PodAntiAffinity](#podantiaffinity-v1)*  | Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).

