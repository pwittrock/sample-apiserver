## PodsMetricStatus v2alpha1 autoscaling

Group        | Version     | Kind
------------ | ---------- | -----------
autoscaling | v2alpha1 | PodsMetricStatus



PodsMetricStatus indicates the current value of a metric describing each pod in the current scale target (for example, transactions-processed-per-second).

<aside class="notice">
Appears In  <a href="#metricstatus-v2alpha1">MetricStatus</a> </aside>

Field        | Description
------------ | -----------
currentAverageValue <br /> *[Quantity](#quantity-resource)*  | currentAverageValue is the current value of the average of the metric across all relevant pods (as a quantity)
metricName <br /> *string*  | metricName is the name of the metric in question

