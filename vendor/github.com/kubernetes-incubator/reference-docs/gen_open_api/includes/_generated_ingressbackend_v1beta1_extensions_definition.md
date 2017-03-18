## IngressBackend v1beta1 extensions

Group        | Version     | Kind
------------ | ---------- | -----------
extensions | v1beta1 | IngressBackend



IngressBackend describes all endpoints for a given service and port.

<aside class="notice">
Appears In  <a href="#httpingresspath-v1beta1">HTTPIngressPath</a>  <a href="#ingressspec-v1beta1">IngressSpec</a> </aside>

Field        | Description
------------ | -----------
serviceName <br /> *string*  | Specifies the name of the referenced service.
servicePort  | Specifies the port of the referenced service.

