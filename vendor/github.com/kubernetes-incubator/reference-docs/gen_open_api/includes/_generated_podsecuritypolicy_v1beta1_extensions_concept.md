

-----------
# PodSecurityPolicy v1beta1 extensions



Group        | Version     | Kind
------------ | ---------- | -----------
extensions | v1beta1 | PodSecurityPolicy







Pod Security Policy governs the ability to make requests that affect the Security Context that will be applied to a pod and container.

<aside class="notice">
Appears In <a href="#podsecuritypolicylist-v1beta1">PodSecurityPolicyList</a> </aside>

Field        | Description
------------ | -----------
apiVersion <br /> *string*  | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#resources
kind <br /> *string*  | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds
metadata <br /> *[ObjectMeta](#objectmeta-v1)*  | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
spec <br /> *[PodSecurityPolicySpec](#podsecuritypolicyspec-v1beta1)*  | spec defines the policy enforced.


### PodSecurityPolicySpec v1beta1

<aside class="notice">
Appears In <a href="#podsecuritypolicy-v1beta1">PodSecurityPolicy</a> </aside>

Field        | Description
------------ | -----------
allowedCapabilities <br /> *string array*  | AllowedCapabilities is a list of capabilities that can be requested to add to the container. Capabilities in this field may be added at the pod author's discretion. You must not list a capability in both AllowedCapabilities and RequiredDropCapabilities.
defaultAddCapabilities <br /> *string array*  | DefaultAddCapabilities is the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability.  You may not list a capabiility in both DefaultAddCapabilities and RequiredDropCapabilities.
fsGroup <br /> *[FSGroupStrategyOptions](#fsgroupstrategyoptions-v1beta1)*  | FSGroup is the strategy that will dictate what fs group is used by the SecurityContext.
hostIPC <br /> *boolean*  | hostIPC determines if the policy allows the use of HostIPC in the pod spec.
hostNetwork <br /> *boolean*  | hostNetwork determines if the policy allows the use of HostNetwork in the pod spec.
hostPID <br /> *boolean*  | hostPID determines if the policy allows the use of HostPID in the pod spec.
hostPorts <br /> *[HostPortRange](#hostportrange-v1beta1) array*  | hostPorts determines which host port ranges are allowed to be exposed.
privileged <br /> *boolean*  | privileged determines if a pod can request to be run as privileged.
readOnlyRootFilesystem <br /> *boolean*  | ReadOnlyRootFilesystem when set to true will force containers to run with a read only root file system.  If the container specifically requests to run with a non-read only root file system the PSP should deny the pod. If set to false the container may run with a read only root file system if it wishes but it will not be forced to.
requiredDropCapabilities <br /> *string array*  | RequiredDropCapabilities are the capabilities that will be dropped from the container.  These are required to be dropped and cannot be added.
runAsUser <br /> *[RunAsUserStrategyOptions](#runasuserstrategyoptions-v1beta1)*  | runAsUser is the strategy that will dictate the allowable RunAsUser values that may be set.
seLinux <br /> *[SELinuxStrategyOptions](#selinuxstrategyoptions-v1beta1)*  | seLinux is the strategy that will dictate the allowable labels that may be set.
supplementalGroups <br /> *[SupplementalGroupsStrategyOptions](#supplementalgroupsstrategyoptions-v1beta1)*  | SupplementalGroups is the strategy that will dictate what supplemental groups are used by the SecurityContext.
volumes <br /> *string array*  | volumes is a white list of allowed volume plugins.  Empty indicates that all plugins may be used.

### PodSecurityPolicyList v1beta1



Field        | Description
------------ | -----------
apiVersion <br /> *string*  | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#resources
items <br /> *[PodSecurityPolicy](#podsecuritypolicy-v1beta1) array*  | Items is a list of schema objects.
kind <br /> *string*  | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds
metadata <br /> *[ListMeta](#listmeta-v1)*  | Standard list metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata




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



create a PodSecurityPolicy

### HTTP Request

`POST /apis/extensions/v1beta1/podsecuritypolicies`


### Query Parameters

Parameter    | Description
------------ | -----------
pretty  | If 'true', then the output is pretty printed.

### Body Parameters

Parameter    | Description
------------ | -----------
body <br /> *[PodSecurityPolicy](#podsecuritypolicy-v1beta1)*  | 

### Response

Code         | Description
------------ | -----------
200 <br /> *[PodSecurityPolicy](#podsecuritypolicy-v1beta1)*  | OK


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



replace the specified PodSecurityPolicy

### HTTP Request

`PUT /apis/extensions/v1beta1/podsecuritypolicies/{name}`

### Path Parameters

Parameter    | Description
------------ | -----------
name  | name of the PodSecurityPolicy

### Query Parameters

Parameter    | Description
------------ | -----------
pretty  | If 'true', then the output is pretty printed.

### Body Parameters

Parameter    | Description
------------ | -----------
body <br /> *[PodSecurityPolicy](#podsecuritypolicy-v1beta1)*  | 

### Response

Code         | Description
------------ | -----------
200 <br /> *[PodSecurityPolicy](#podsecuritypolicy-v1beta1)*  | OK


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



partially update the specified PodSecurityPolicy

### HTTP Request

`PATCH /apis/extensions/v1beta1/podsecuritypolicies/{name}`

### Path Parameters

Parameter    | Description
------------ | -----------
name  | name of the PodSecurityPolicy

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
200 <br /> *[PodSecurityPolicy](#podsecuritypolicy-v1beta1)*  | OK


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



delete a PodSecurityPolicy

### HTTP Request

`DELETE /apis/extensions/v1beta1/podsecuritypolicies/{name}`

### Path Parameters

Parameter    | Description
------------ | -----------
name  | name of the PodSecurityPolicy

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



delete collection of PodSecurityPolicy

### HTTP Request

`DELETE /apis/extensions/v1beta1/podsecuritypolicies`


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



read the specified PodSecurityPolicy

### HTTP Request

`GET /apis/extensions/v1beta1/podsecuritypolicies/{name}`

### Path Parameters

Parameter    | Description
------------ | -----------
name  | name of the PodSecurityPolicy

### Query Parameters

Parameter    | Description
------------ | -----------
pretty  | If 'true', then the output is pretty printed.
exact  | Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'.
export  | Should this value be exported.  Export strips fields that a user can not specify.


### Response

Code         | Description
------------ | -----------
200 <br /> *[PodSecurityPolicy](#podsecuritypolicy-v1beta1)*  | OK


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



list or watch objects of kind PodSecurityPolicy

### HTTP Request

`GET /apis/extensions/v1beta1/podsecuritypolicies`


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
200 <br /> *[PodSecurityPolicyList](#podsecuritypolicylist-v1beta1)*  | OK


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



watch changes to an object of kind PodSecurityPolicy

### HTTP Request

`GET /apis/extensions/v1beta1/watch/podsecuritypolicies/{name}`

### Path Parameters

Parameter    | Description
------------ | -----------
name  | name of the PodSecurityPolicy

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



watch individual changes to a list of PodSecurityPolicy

### HTTP Request

`GET /apis/extensions/v1beta1/watch/podsecuritypolicies`


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




