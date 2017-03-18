## PodAffinityTerm v1 core

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | PodAffinityTerm



Defines a set of pods (namely those matching the labelSelector relative to the given namespace(s)) that this pod should be co-located (affinity) or not co-located (anti-affinity) with, where co-located is defined as running on a node whose value of the label with key <topologyKey> tches that of any node on which a pod of the set of pods is running

<aside class="notice">
Appears In  <a href="#podaffinity-v1">PodAffinity</a>  <a href="#podantiaffinity-v1">PodAntiAffinity</a>  <a href="#weightedpodaffinityterm-v1">WeightedPodAffinityTerm</a> </aside>

Field        | Description
------------ | -----------
labelSelector <br /> *[LabelSelector](#labelselector-v1)*  | A label query over a set of resources, in this case pods.
namespaces <br /> *string array*  | namespaces specifies which namespaces the labelSelector applies to (matches against); nil list means "this pod's namespace," empty list means "all namespaces" The json tag here is not "omitempty" since we need to distinguish nil and empty. See https://golang.org/pkg/encoding/json/#Marshal for more details.
topologyKey <br /> *string*  | This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. For PreferredDuringScheduling pod anti-affinity, empty topologyKey is interpreted as "all topologies" ("all topologies" here means all the topologyKeys indicated by scheduler command-line argument --failure-domains); for affinity and for RequiredDuringScheduling pod anti-affinity, empty topologyKey is not allowed.

