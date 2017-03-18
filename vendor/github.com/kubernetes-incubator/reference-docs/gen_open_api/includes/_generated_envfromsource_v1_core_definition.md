## EnvFromSource v1 core

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | EnvFromSource



EnvFromSource represents the source of a set of ConfigMaps

<aside class="notice">
Appears In  <a href="#container-v1">Container</a>  <a href="#podpresetspec-v1alpha1">PodPresetSpec</a> </aside>

Field        | Description
------------ | -----------
configMapRef <br /> *[ConfigMapEnvSource](#configmapenvsource-v1)*  | The ConfigMap to select from
prefix <br /> *string*  | An optional identifer to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
secretRef <br /> *[SecretEnvSource](#secretenvsource-v1)*  | The Secret to select from

