## PortworxVolumeSource v1 core

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | PortworxVolumeSource



PortworxVolumeSource represents a Portworx volume resource.

<aside class="notice">
Appears In  <a href="#persistentvolumespec-v1">PersistentVolumeSpec</a>  <a href="#volume-v1">Volume</a> </aside>

Field        | Description
------------ | -----------
fsType <br /> *string*  | FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs". Implicitly inferred to be "ext4" if unspecified.
readOnly <br /> *boolean*  | Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
volumeID <br /> *string*  | VolumeID uniquely identifies a Portworx volume

