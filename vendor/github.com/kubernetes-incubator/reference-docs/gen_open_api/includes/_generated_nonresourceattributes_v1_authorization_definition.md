## NonResourceAttributes v1 authorization

Group        | Version     | Kind
------------ | ---------- | -----------
authorization | v1 | NonResourceAttributes

<aside class="notice">Other api versions of this object exist: <a href="#nonresourceattributes-v1beta1">v1beta1</a> </aside>

NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface

<aside class="notice">
Appears In  <a href="#selfsubjectaccessreviewspec-v1">SelfSubjectAccessReviewSpec</a>  <a href="#subjectaccessreviewspec-v1">SubjectAccessReviewSpec</a> </aside>

Field        | Description
------------ | -----------
path <br /> *string*  | Path is the URL path of the request
verb <br /> *string*  | Verb is the standard HTTP verb

