

-----------
# RoleBinding v1beta1 rbac



Group        | Version     | Kind
------------ | ---------- | -----------
rbac | v1beta1 | RoleBinding




<aside class="notice">Other api versions of this object exist: <a href="#rolebinding-v1alpha1">v1alpha1</a> </aside>


RoleBinding references a role, but does not contain it.  It can reference a Role in the same namespace or a ClusterRole in the global namespace. It adds who information via Subjects and namespace information by which namespace it exists in.  RoleBindings in a given namespace only have effect in that namespace.

<aside class="notice">
Appears In <a href="#rolebindinglist-v1beta1">RoleBindingList</a> </aside>

Field        | Description
------------ | -----------
apiVersion <br /> *string*  | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#resources
kind <br /> *string*  | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds
metadata <br /> *[ObjectMeta](#objectmeta-v1)*  | Standard object's metadata.
roleRef <br /> *[RoleRef](#roleref-v1beta1)*  | RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
subjects <br /> *[Subject](#subject-v1beta1) array*  | Subjects holds references to the objects the role applies to.


### RoleBindingList v1beta1



Field        | Description
------------ | -----------
apiVersion <br /> *string*  | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#resources
items <br /> *[RoleBinding](#rolebinding-v1beta1) array*  | Items is a list of RoleBindings
kind <br /> *string*  | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds
metadata <br /> *[ListMeta](#listmeta-v1)*  | Standard object's metadata.




## <strong>Write Operations</strong>

See supported operations below...

## Create

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



create a RoleBinding

### HTTP Request

`POST /apis/rbac.authorization.k8s.io/v1beta1/namespaces/{namespace}/rolebindings`

### Path Parameters

Parameter    | Description
------------ | -----------
namespace  | object name and auth scope, such as for teams and projects

### Query Parameters

Parameter    | Description
------------ | -----------
pretty  | If 'true', then the output is pretty printed.

### Body Parameters

Parameter    | Description
------------ | -----------
body <br /> *[RoleBinding](#rolebinding-v1beta1)*  | 

### Response

Code         | Description
------------ | -----------
200 <br /> *[RoleBinding](#rolebinding-v1beta1)*  | OK


## Replace

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



replace the specified RoleBinding

### HTTP Request

`PUT /apis/rbac.authorization.k8s.io/v1beta1/namespaces/{namespace}/rolebindings/{name}`

### Path Parameters

Parameter    | Description
------------ | -----------
name  | name of the RoleBinding
namespace  | object name and auth scope, such as for teams and projects

### Query Parameters

Parameter    | Description
------------ | -----------
pretty  | If 'true', then the output is pretty printed.

### Body Parameters

Parameter    | Description
------------ | -----------
body <br /> *[RoleBinding](#rolebinding-v1beta1)*  | 

### Response

Code         | Description
------------ | -----------
200 <br /> *[RoleBinding](#rolebinding-v1beta1)*  | OK


## Patch

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



partially update the specified RoleBinding

### HTTP Request

`PATCH /apis/rbac.authorization.k8s.io/v1beta1/namespaces/{namespace}/rolebindings/{name}`

### Path Parameters

Parameter    | Description
------------ | -----------
name  | name of the RoleBinding
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
200 <br /> *[RoleBinding](#rolebinding-v1beta1)*  | OK


## Delete

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



delete a RoleBinding

### HTTP Request

`DELETE /apis/rbac.authorization.k8s.io/v1beta1/namespaces/{namespace}/rolebindings/{name}`

### Path Parameters

Parameter    | Description
------------ | -----------
name  | name of the RoleBinding
namespace  | object name and auth scope, such as for teams and projects

### Query Parameters

Parameter    | Description
------------ | -----------
pretty  | If 'true', then the output is pretty printed.
gracePeriodSeconds  | The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.
orphanDependents  | Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the "orphan" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both.
propagationPolicy  | Whether and how garbage collection will be performed. Defaults to Default. Either this field or OrphanDependents may be set, but not both.

### Body Parameters

Parameter    | Description
------------ | -----------
body <br /> *[DeleteOptions](#deleteoptions-v1)*  | 

### Response

Code         | Description
------------ | -----------
200 <br /> *[Status](#status-v1)*  | OK


## Delete Collection

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



delete collection of RoleBinding

### HTTP Request

`DELETE /apis/rbac.authorization.k8s.io/v1beta1/namespaces/{namespace}/rolebindings`

### Path Parameters

Parameter    | Description
------------ | -----------
namespace  | object name and auth scope, such as for teams and projects

### Query Parameters

Parameter    | Description
------------ | -----------
pretty  | If 'true', then the output is pretty printed.
fieldSelector  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
resourceVersion  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.
timeoutSeconds  | Timeout for the list/watch call.
watch  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.


### Response

Code         | Description
------------ | -----------
200 <br /> *[Status](#status-v1)*  | OK



## <strong>Read Operations</strong>

See supported operations below...

## Read

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



read the specified RoleBinding

### HTTP Request

`GET /apis/rbac.authorization.k8s.io/v1beta1/namespaces/{namespace}/rolebindings/{name}`

### Path Parameters

Parameter    | Description
------------ | -----------
name  | name of the RoleBinding
namespace  | object name and auth scope, such as for teams and projects

### Query Parameters

Parameter    | Description
------------ | -----------
pretty  | If 'true', then the output is pretty printed.


### Response

Code         | Description
------------ | -----------
200 <br /> *[RoleBinding](#rolebinding-v1beta1)*  | OK


## List

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



list or watch objects of kind RoleBinding

### HTTP Request

`GET /apis/rbac.authorization.k8s.io/v1beta1/namespaces/{namespace}/rolebindings`

### Path Parameters

Parameter    | Description
------------ | -----------
namespace  | object name and auth scope, such as for teams and projects

### Query Parameters

Parameter    | Description
------------ | -----------
pretty  | If 'true', then the output is pretty printed.
fieldSelector  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
resourceVersion  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.
timeoutSeconds  | Timeout for the list/watch call.
watch  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.


### Response

Code         | Description
------------ | -----------
200 <br /> *[RoleBindingList](#rolebindinglist-v1beta1)*  | OK


## List All Namespaces

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



list or watch objects of kind RoleBinding

### HTTP Request

`GET /apis/rbac.authorization.k8s.io/v1beta1/rolebindings`


### Query Parameters

Parameter    | Description
------------ | -----------
fieldSelector  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
pretty  | If 'true', then the output is pretty printed.
resourceVersion  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.
timeoutSeconds  | Timeout for the list/watch call.
watch  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.


### Response

Code         | Description
------------ | -----------
200 <br /> *[RoleBindingList](#rolebindinglist-v1beta1)*  | OK


## Watch

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



watch changes to an object of kind RoleBinding

### HTTP Request

`GET /apis/rbac.authorization.k8s.io/v1beta1/watch/namespaces/{namespace}/rolebindings/{name}`

### Path Parameters

Parameter    | Description
------------ | -----------
name  | name of the RoleBinding
namespace  | object name and auth scope, such as for teams and projects

### Query Parameters

Parameter    | Description
------------ | -----------
fieldSelector  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
pretty  | If 'true', then the output is pretty printed.
resourceVersion  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.
timeoutSeconds  | Timeout for the list/watch call.
watch  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.


### Response

Code         | Description
------------ | -----------
200 <br /> *[WatchEvent](#watchevent-v1)*  | OK


## Watch List

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



watch individual changes to a list of RoleBinding

### HTTP Request

`GET /apis/rbac.authorization.k8s.io/v1beta1/watch/namespaces/{namespace}/rolebindings`

### Path Parameters

Parameter    | Description
------------ | -----------
namespace  | object name and auth scope, such as for teams and projects

### Query Parameters

Parameter    | Description
------------ | -----------
fieldSelector  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
pretty  | If 'true', then the output is pretty printed.
resourceVersion  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.
timeoutSeconds  | Timeout for the list/watch call.
watch  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.


### Response

Code         | Description
------------ | -----------
200 <br /> *[WatchEvent](#watchevent-v1)*  | OK


## Watch List All Namespaces

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



watch individual changes to a list of RoleBinding

### HTTP Request

`GET /apis/rbac.authorization.k8s.io/v1beta1/watch/rolebindings`


### Query Parameters

Parameter    | Description
------------ | -----------
fieldSelector  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
pretty  | If 'true', then the output is pretty printed.
resourceVersion  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.
timeoutSeconds  | Timeout for the list/watch call.
watch  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.


### Response

Code         | Description
------------ | -----------
200 <br /> *[WatchEvent](#watchevent-v1)*  | OK



