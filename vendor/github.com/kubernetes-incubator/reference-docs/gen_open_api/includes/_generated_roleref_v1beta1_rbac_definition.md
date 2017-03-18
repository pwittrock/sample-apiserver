## RoleRef v1beta1 rbac

Group        | Version     | Kind
------------ | ---------- | -----------
rbac | v1beta1 | RoleRef

<aside class="notice">Other api versions of this object exist: <a href="#roleref-v1alpha1">v1alpha1</a> </aside>

RoleRef contains information that points to the role being used

<aside class="notice">
Appears In  <a href="#clusterrolebinding-v1beta1">ClusterRoleBinding</a>  <a href="#rolebinding-v1beta1">RoleBinding</a> </aside>

Field        | Description
------------ | -----------
apiGroup <br /> *string*  | APIGroup is the group for the resource being referenced
kind <br /> *string*  | Kind is the type of resource being referenced
name <br /> *string*  | Name is the name of resource being referenced

