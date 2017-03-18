

-----------
# HorizontalPodAutoscalerSpec v2alpha1 autoscaling



Group        | Version     | Kind
------------ | ---------- | -----------
autoscaling | v2alpha1 | HorizontalPodAutoscalerSpec




<aside class="notice">Other api versions of this object exist: <a href="#horizontalpodautoscalerspec-v1">v1</a> </aside>


HorizontalPodAutoscalerSpec describes the desired functionality of the HorizontalPodAutoscaler.

<aside class="notice">
Appears In <a href="#horizontalpodautoscaler-v2alpha1">HorizontalPodAutoscaler</a> </aside>

Field        | Description
------------ | -----------
maxReplicas <br /> *integer*  | maxReplicas is the upper limit for the number of replicas to which the autoscaler can scale up. It cannot be less that minReplicas.
metrics <br /> *[MetricSpec](#metricspec-v2alpha1) array*  | metrics contains the specifications for which to use to calculate the desired replica count (the maximum replica count across all metrics will be used).  The desired replica count is calculated multiplying the ratio between the target value and the current value by the current number of pods.  Ergo, metrics used must decrease as the pod count is increased, and vice-versa.  See the individual metric source types for more information about how each type of metric must respond.
minReplicas <br /> *integer*  | minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down. It defaults to 1 pod.
scaleTargetRef <br /> *[CrossVersionObjectReference](#crossversionobjectreference-v2alpha1)*  | scaleTargetRef points to the target resource to scale, and is used to the pods for which metrics should be collected, as well as to actually change the replica count.






