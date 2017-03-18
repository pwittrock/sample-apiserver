## ObjectMetricStatus v2alpha1 autoscaling

Group        | Version     | Kind
------------ | ---------- | -----------
autoscaling | v2alpha1 | ObjectMetricStatus



ObjectMetricStatus indicates the current value of a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).

<aside class="notice">
Appears In  <a href="#metricstatus-v2alpha1">MetricStatus</a> </aside>

Field        | Description
------------ | -----------
currentValue <br /> *[Quantity](#quantity-resource)*  | currentValue is the current value of the metric (as a quantity).
metricName <br /> *string*  | metricName is the name of the metric in question.
target <br /> *[CrossVersionObjectReference](#crossversionobjectreference-v2alpha1)*  | target is the described Kubernetes object.

