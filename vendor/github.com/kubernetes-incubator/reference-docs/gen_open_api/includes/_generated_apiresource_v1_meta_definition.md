## APIResource v1 meta

Group        | Version     | Kind
------------ | ---------- | -----------
meta | v1 | APIResource



APIResource specifies the name of a resource and whether it is namespaced.

<aside class="notice">
Appears In  <a href="#apiresourcelist-v1">APIResourceList</a> </aside>

Field        | Description
------------ | -----------
kind <br /> *string*  | kind is the kind for the resource (e.g. 'Foo' is the kind for a resource 'foo')
name <br /> *string*  | name is the name of the resource.
namespaced <br /> *boolean*  | namespaced indicates if a resource is namespaced or not.
shortNames <br /> *string array*  | shortNames is a list of suggested short names of the resource.
verbs <br /> *string array*  | verbs is a list of supported kube verbs (this includes get, list, watch, create, update, patch, delete, deletecollection, and proxy)

