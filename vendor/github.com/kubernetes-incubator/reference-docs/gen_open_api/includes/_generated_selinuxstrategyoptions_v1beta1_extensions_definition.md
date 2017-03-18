## SELinuxStrategyOptions v1beta1 extensions

Group        | Version     | Kind
------------ | ---------- | -----------
extensions | v1beta1 | SELinuxStrategyOptions



SELinux  Strategy Options defines the strategy type and any options used to create the strategy.

<aside class="notice">
Appears In  <a href="#podsecuritypolicyspec-v1beta1">PodSecurityPolicySpec</a> </aside>

Field        | Description
------------ | -----------
rule <br /> *string*  | type is the strategy that will dictate the allowable labels that may be set.
seLinuxOptions <br /> *[SELinuxOptions](#selinuxoptions-v1)*  | seLinuxOptions required to run as; required for MustRunAs More info: http://releases.k8s.io/HEAD/docs/design/security_context.md#security-context

