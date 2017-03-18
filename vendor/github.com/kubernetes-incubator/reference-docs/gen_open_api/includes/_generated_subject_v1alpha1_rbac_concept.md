

-----------
# Subject v1alpha1 rbac



Group        | Version     | Kind
------------ | ---------- | -----------
rbac | v1alpha1 | Subject




<aside class="notice">Other api versions of this object exist: <a href="#subject-v1beta1">v1beta1</a> </aside>


Subject contains a reference to the object or user identities a role binding applies to.  This can either hold a direct API object reference, or a value for non-objects such as user and group names.

<aside class="notice">
Appears In <a href="#clusterrolebinding-v1alpha1">ClusterRoleBinding</a> <a href="#rolebinding-v1alpha1">RoleBinding</a> </aside>

Field        | Description
------------ | -----------
apiVersion <br /> *string*  | APIVersion holds the API group and version of the referenced subject. Defaults to "v1" for ServiceAccount subjects. Defaults to "rbac.authorization.k8s.io/v1alpha1" for User and Group subjects.
kind <br /> *string*  | Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount". If the Authorizer does not recognized the kind value, the Authorizer should report an error.
name <br /> *string*  | Name of the object being referenced.
namespace <br /> *string*  | Namespace of the referenced object.  If the object kind is non-namespace, such as "User" or "Group", and this value is not empty the Authorizer should report an error.






