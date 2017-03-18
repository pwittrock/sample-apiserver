## ObjectMetricSource v2alpha1 autoscaling

Group        | Version     | Kind
------------ | ---------- | -----------
autoscaling | v2alpha1 | ObjectMetricSource



ObjectMetricSource indicates how to scale on a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).

<aside class="notice">
Appears In  <a href="#metricspec-v2alpha1">MetricSpec</a> </aside>

Field        | Description
------------ | -----------
metricName <br /> *string*  | metricName is the name of the metric in question.
target <br /> *[CrossVersionObjectReference](#crossversionobjectreference-v2alpha1)*  | target is the described Kubernetes object.
targetValue <br /> *[Quantity](#quantity-resource)*  | targetValue is the target value of the metric (as a quantity).

