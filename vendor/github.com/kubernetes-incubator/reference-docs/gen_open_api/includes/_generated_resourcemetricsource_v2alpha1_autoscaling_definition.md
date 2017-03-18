## ResourceMetricSource v2alpha1 autoscaling

Group        | Version     | Kind
------------ | ---------- | -----------
autoscaling | v2alpha1 | ResourceMetricSource



ResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  The values will be averaged together before being compared to the target.  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.  Only one "target" type should be set.

<aside class="notice">
Appears In  <a href="#metricspec-v2alpha1">MetricSpec</a> </aside>

Field        | Description
------------ | -----------
name <br /> *string*  | name is the name of the resource in question.
targetAverageUtilization <br /> *integer*  | targetAverageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.
targetAverageValue <br /> *[Quantity](#quantity-resource)*  | targetAverageValue is the the target value of the average of the resource metric across all relevant pods, as a raw value (instead of as a percentage of the request), similar to the "pods" metric source type.

