

-----------
# HorizontalPodAutoscalerStatus v2alpha1 autoscaling



Group        | Version     | Kind
------------ | ---------- | -----------
autoscaling | v2alpha1 | HorizontalPodAutoscalerStatus




<aside class="notice">Other api versions of this object exist: <a href="#horizontalpodautoscalerstatus-v1">v1</a> </aside>


HorizontalPodAutoscalerStatus describes the current status of a horizontal pod autoscaler.

<aside class="notice">
Appears In <a href="#horizontalpodautoscaler-v2alpha1">HorizontalPodAutoscaler</a> </aside>

Field        | Description
------------ | -----------
currentMetrics <br /> *[MetricStatus](#metricstatus-v2alpha1) array*  | currentMetrics is the last read state of the metrics used by this autoscaler.
currentReplicas <br /> *integer*  | currentReplicas is current number of replicas of pods managed by this autoscaler, as last seen by the autoscaler.
desiredReplicas <br /> *integer*  | desiredReplicas is the desired number of replicas of pods managed by this autoscaler, as last calculated by the autoscaler.
lastScaleTime <br /> *[Time](#time-v1)*  | lastScaleTime is the last time the HorizontalPodAutoscaler scaled the number of pods, used by the autoscaler to control how often the number of pods is changed.
observedGeneration <br /> *integer*  | observedGeneration is the most recent generation observed by this autoscaler.






