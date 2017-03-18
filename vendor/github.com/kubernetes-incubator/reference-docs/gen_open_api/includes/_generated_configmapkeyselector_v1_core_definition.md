## ConfigMapKeySelector v1 core

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | ConfigMapKeySelector



Selects a key from a ConfigMap.

<aside class="notice">
Appears In  <a href="#envvarsource-v1">EnvVarSource</a> </aside>

Field        | Description
------------ | -----------
key <br /> *string*  | The key to select.
name <br /> *string*  | Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
optional <br /> *boolean*  | Specify whether the ConfigMap or it's key must be defined

