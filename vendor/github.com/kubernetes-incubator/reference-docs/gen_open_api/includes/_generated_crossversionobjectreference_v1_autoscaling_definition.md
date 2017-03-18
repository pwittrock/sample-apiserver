## CrossVersionObjectReference v1 autoscaling

Group        | Version     | Kind
------------ | ---------- | -----------
autoscaling | v1 | CrossVersionObjectReference

<aside class="notice">Other api versions of this object exist: <a href="#crossversionobjectreference-v2alpha1">v2alpha1</a> </aside>

CrossVersionObjectReference contains enough information to let you identify the referred resource.

<aside class="notice">
Appears In  <a href="#horizontalpodautoscalerspec-v1">HorizontalPodAutoscalerSpec</a> </aside>

Field        | Description
------------ | -----------
apiVersion <br /> *string*  | API version of the referent
kind <br /> *string*  | Kind of the referent; More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds"
name <br /> *string*  | Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names

