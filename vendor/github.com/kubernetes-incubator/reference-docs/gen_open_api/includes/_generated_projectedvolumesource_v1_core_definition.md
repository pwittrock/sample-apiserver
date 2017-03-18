## ProjectedVolumeSource v1 core

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | ProjectedVolumeSource



Represents a projected volume source

<aside class="notice">
Appears In  <a href="#volume-v1">Volume</a> </aside>

Field        | Description
------------ | -----------
defaultMode <br /> *integer*  | Mode bits to use on created files by default. Must be a value between 0 and 0777. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
sources <br /> *[VolumeProjection](#volumeprojection-v1) array*  | list of volume projections

