## GroupVersionForDiscovery v1 meta

Group        | Version     | Kind
------------ | ---------- | -----------
meta | v1 | GroupVersionForDiscovery



GroupVersion contains the "group/version" and "version" string of a version. It is made a struct to keep extensibility.

<aside class="notice">
Appears In  <a href="#apigroup-v1">APIGroup</a> </aside>

Field        | Description
------------ | -----------
groupVersion <br /> *string*  | groupVersion specifies the API group and version in the form "group/version"
version <br /> *string*  | version specifies the version in the form of "version". This is to save the clients the trouble of splitting the GroupVersion.

