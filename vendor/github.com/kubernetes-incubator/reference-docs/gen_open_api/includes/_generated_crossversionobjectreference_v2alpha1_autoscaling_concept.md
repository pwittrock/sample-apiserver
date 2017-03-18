

-----------
# CrossVersionObjectReference v2alpha1 autoscaling



Group        | Version     | Kind
------------ | ---------- | -----------
autoscaling | v2alpha1 | CrossVersionObjectReference




<aside class="notice">Other api versions of this object exist: <a href="#crossversionobjectreference-v1">v1</a> </aside>


CrossVersionObjectReference contains enough information to let you identify the referred resource.

<aside class="notice">
Appears In <a href="#horizontalpodautoscalerspec-v2alpha1">HorizontalPodAutoscalerSpec</a> <a href="#objectmetricsource-v2alpha1">ObjectMetricSource</a> <a href="#objectmetricstatus-v2alpha1">ObjectMetricStatus</a> </aside>

Field        | Description
------------ | -----------
apiVersion <br /> *string*  | API version of the referent
kind <br /> *string*  | Kind of the referent; More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds"
name <br /> *string*  | Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names






