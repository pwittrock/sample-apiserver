

-----------
# Scale v1beta1 apps



Group        | Version     | Kind
------------ | ---------- | -----------
apps | v1beta1 | Scale




<aside class="notice">Other api versions of this object exist: <a href="#scale-v1">v1</a> </aside>


Scale represents a scaling request for a resource.



Field        | Description
------------ | -----------
apiVersion <br /> *string*  | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#resources
kind <br /> *string*  | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds
metadata <br /> *[ObjectMeta](#objectmeta-v1)*  | Standard object metadata; More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata.
spec <br /> *[ScaleSpec](#scalespec-v1beta1)*  | defines the behavior of the scale. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status.
status <br /> *[ScaleStatus](#scalestatus-v1beta1)*  | current status of the scale. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status. Read-only.


### ScaleSpec v1beta1

<aside class="notice">
Appears In <a href="#scale-v1beta1">Scale</a> </aside>

Field        | Description
------------ | -----------
replicas <br /> *integer*  | desired number of instances for the scaled object.

### ScaleStatus v1beta1

<aside class="notice">
Appears In <a href="#scale-v1beta1">Scale</a> </aside>

Field        | Description
------------ | -----------
replicas <br /> *integer*  | actual number of observed instances of the scaled object.
selector <br /> *object*  | label query over pods that should match the replicas count. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors
targetSelector <br /> *string*  | label selector for pods that should match the replicas count. This is a serializated version of both map-based and more expressive set-based selectors. This is done to avoid introspection in the clients. The string will be in the same format as the query-param syntax. If the target type only supports map-based selectors, both this field and map-based selector field are populated. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors




## <strong>Misc Operations</strong>

See supported operations below...

## Read Scale

>bdocs-tab:kubectl `kubectl` Command

```bdocs-tab:kubectl_shell

Coming Soon

```

>bdocs-tab:curl `curl` Command (*requires `kubectl proxy` to be running*)

```bdocs-tab:curl_shell

Coming Soon

```

>bdocs-tab:kubectl Output

```bdocs-tab:kubectl_json

Coming Soon

```
>bdocs-tab:curl Response Body

```bdocs-tab:curl_json

Coming Soon

```



read scale of the specified Scale

### HTTP Request

`GET /apis/apps/v1beta1/namespaces/{namespace}/deployments/{name}/scale`

### Path Parameters

Parameter    | Description
------------ | -----------
name  | name of the Scale
namespace  | object name and auth scope, such as for teams and projects

### Query Parameters

Parameter    | Description
------------ | -----------
pretty  | If 'true', then the output is pretty printed.


### Response

Code         | Description
------------ | -----------
200 <br /> *[Scale](#scale-v1beta1)*  | OK


## Replace Scale

>bdocs-tab:kubectl `kubectl` Command

```bdocs-tab:kubectl_shell

Coming Soon

```

>bdocs-tab:curl `curl` Command (*requires `kubectl proxy` to be running*)

```bdocs-tab:curl_shell

Coming Soon

```

>bdocs-tab:kubectl Output

```bdocs-tab:kubectl_json

Coming Soon

```
>bdocs-tab:curl Response Body

```bdocs-tab:curl_json

Coming Soon

```



replace scale of the specified Scale

### HTTP Request

`PUT /apis/apps/v1beta1/namespaces/{namespace}/deployments/{name}/scale`

### Path Parameters

Parameter    | Description
------------ | -----------
name  | name of the Scale
namespace  | object name and auth scope, such as for teams and projects

### Query Parameters

Parameter    | Description
------------ | -----------
pretty  | If 'true', then the output is pretty printed.

### Body Parameters

Parameter    | Description
------------ | -----------
body <br /> *[Scale](#scale-v1beta1)*  | 

### Response

Code         | Description
------------ | -----------
200 <br /> *[Scale](#scale-v1beta1)*  | OK


## Patch Scale

>bdocs-tab:kubectl `kubectl` Command

```bdocs-tab:kubectl_shell

Coming Soon

```

>bdocs-tab:curl `curl` Command (*requires `kubectl proxy` to be running*)

```bdocs-tab:curl_shell

Coming Soon

```

>bdocs-tab:kubectl Output

```bdocs-tab:kubectl_json

Coming Soon

```
>bdocs-tab:curl Response Body

```bdocs-tab:curl_json

Coming Soon

```



partially update scale of the specified Scale

### HTTP Request

`PATCH /apis/apps/v1beta1/namespaces/{namespace}/deployments/{name}/scale`

### Path Parameters

Parameter    | Description
------------ | -----------
name  | name of the Scale
namespace  | object name and auth scope, such as for teams and projects

### Query Parameters

Parameter    | Description
------------ | -----------
pretty  | If 'true', then the output is pretty printed.

### Body Parameters

Parameter    | Description
------------ | -----------
body <br /> *[Patch](#patch-v1)*  | 

### Response

Code         | Description
------------ | -----------
200 <br /> *[Scale](#scale-v1beta1)*  | OK




