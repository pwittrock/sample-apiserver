## MetricStatus v2alpha1 autoscaling

Group        | Version     | Kind
------------ | ---------- | -----------
autoscaling | v2alpha1 | MetricStatus



MetricStatus describes the last-read state of a single metric.

<aside class="notice">
Appears In  <a href="#horizontalpodautoscalerstatus-v2alpha1">HorizontalPodAutoscalerStatus</a> </aside>

Field        | Description
------------ | -----------
object <br /> *[ObjectMetricStatus](#objectmetricstatus-v2alpha1)*  | object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).
pods <br /> *[PodsMetricStatus](#podsmetricstatus-v2alpha1)*  | pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value.
resource <br /> *[ResourceMetricStatus](#resourcemetricstatus-v2alpha1)*  | resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
type <br /> *string*  | type is the type of metric source.  It will match one of the fields below.

