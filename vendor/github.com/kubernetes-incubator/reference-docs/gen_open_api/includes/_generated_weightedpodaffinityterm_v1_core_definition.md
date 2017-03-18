## WeightedPodAffinityTerm v1 core

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | WeightedPodAffinityTerm



The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)

<aside class="notice">
Appears In  <a href="#podaffinity-v1">PodAffinity</a>  <a href="#podantiaffinity-v1">PodAntiAffinity</a> </aside>

Field        | Description
------------ | -----------
podAffinityTerm <br /> *[PodAffinityTerm](#podaffinityterm-v1)*  | Required. A pod affinity term, associated with the corresponding weight.
weight <br /> *integer*  | weight associated with matching the corresponding podAffinityTerm, in the range 1-100.

