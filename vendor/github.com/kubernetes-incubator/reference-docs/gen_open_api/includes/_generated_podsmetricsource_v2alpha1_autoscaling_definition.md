## PodsMetricSource v2alpha1 autoscaling

Group        | Version     | Kind
------------ | ---------- | -----------
autoscaling | v2alpha1 | PodsMetricSource



PodsMetricSource indicates how to scale on a metric describing each pod in the current scale target (for example, transactions-processed-per-second). The values will be averaged together before being compared to the target value.

<aside class="notice">
Appears In  <a href="#metricspec-v2alpha1">MetricSpec</a> </aside>

Field        | Description
------------ | -----------
metricName <br /> *string*  | metricName is the name of the metric in question
targetAverageValue <br /> *[Quantity](#quantity-resource)*  | targetAverageValue is the target value of the average of the metric across all relevant pods (as a quantity)

