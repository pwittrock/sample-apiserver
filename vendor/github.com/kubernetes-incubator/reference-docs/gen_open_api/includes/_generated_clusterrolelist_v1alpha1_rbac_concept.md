

-----------
# ClusterRoleList v1alpha1 rbac



Group        | Version     | Kind
------------ | ---------- | -----------
rbac | v1alpha1 | ClusterRoleList




<aside class="notice">Other api versions of this object exist: <a href="#clusterrolelist-v1beta1">v1beta1</a> </aside>


ClusterRoleList is a collection of ClusterRoles



Field        | Description
------------ | -----------
apiVersion <br /> *string*  | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#resources
items <br /> *[ClusterRole](#clusterrole-v1alpha1) array*  | Items is a list of ClusterRoles
kind <br /> *string*  | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds
metadata <br /> *[ListMeta](#listmeta-v1)*  | Standard object's metadata.






