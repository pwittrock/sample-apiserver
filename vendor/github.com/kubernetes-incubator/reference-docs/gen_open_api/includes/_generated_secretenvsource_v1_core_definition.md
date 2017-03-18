## SecretEnvSource v1 core

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | SecretEnvSource



SecretEnvSource selects a Secret to populate the environment variables with.

The contents of the target Secret's Data field will represent the key-value pairs as environment variables.

<aside class="notice">
Appears In  <a href="#envfromsource-v1">EnvFromSource</a> </aside>

Field        | Description
------------ | -----------
name <br /> *string*  | Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
optional <br /> *boolean*  | Specify whether the Secret must be defined

