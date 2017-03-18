## ConfigMapProjection v1 core

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | ConfigMapProjection



Adapts a ConfigMap into a projected volume.

The contents of the target ConfigMap's Data field will be presented in a projected volume as files using the keys in the Data field as the file names, unless the items element is populated with specific mappings of keys to paths. Note that this is identical to a configmap volume source without the default mode.

<aside class="notice">
Appears In  <a href="#volumeprojection-v1">VolumeProjection</a> </aside>

Field        | Description
------------ | -----------
items <br /> *[KeyToPath](#keytopath-v1) array*  | If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.
name <br /> *string*  | Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
optional <br /> *boolean*  | Specify whether the ConfigMap or it's keys must be defined

