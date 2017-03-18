## ScaleIOVolumeSource v1 core

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | ScaleIOVolumeSource



ScaleIOVolumeSource represents a persistent ScaleIO volume

<aside class="notice">
Appears In  <a href="#persistentvolumespec-v1">PersistentVolumeSpec</a>  <a href="#volume-v1">Volume</a> </aside>

Field        | Description
------------ | -----------
fsType <br /> *string*  | Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
gateway <br /> *string*  | The host address of the ScaleIO API Gateway.
protectionDomain <br /> *string*  | The name of the Protection Domain for the configured storage (defaults to "default").
readOnly <br /> *boolean*  | Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
secretRef <br /> *[LocalObjectReference](#localobjectreference-v1)*  | SecretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.
sslEnabled <br /> *boolean*  | Flag to enable/disable SSL communication with Gateway, default false
storageMode <br /> *string*  | Indicates whether the storage for a volume should be thick or thin (defaults to "thin").
storagePool <br /> *string*  | The Storage Pool associated with the protection domain (defaults to "default").
system <br /> *string*  | The name of the storage system as configured in ScaleIO.
volumeName <br /> *string*  | The name of a volume already created in the ScaleIO system that is associated with this volume source.

