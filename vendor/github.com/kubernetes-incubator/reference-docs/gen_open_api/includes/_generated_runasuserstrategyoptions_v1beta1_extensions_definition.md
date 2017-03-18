## RunAsUserStrategyOptions v1beta1 extensions

Group        | Version     | Kind
------------ | ---------- | -----------
extensions | v1beta1 | RunAsUserStrategyOptions



Run A sUser Strategy Options defines the strategy type and any options used to create the strategy.

<aside class="notice">
Appears In  <a href="#podsecuritypolicyspec-v1beta1">PodSecurityPolicySpec</a> </aside>

Field        | Description
------------ | -----------
ranges <br /> *[IDRange](#idrange-v1beta1) array*  | Ranges are the allowed ranges of uids that may be used.
rule <br /> *string*  | Rule is the strategy that will dictate the allowable RunAsUser values that may be set.

