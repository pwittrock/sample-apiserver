## SupplementalGroupsStrategyOptions v1beta1 extensions

Group        | Version     | Kind
------------ | ---------- | -----------
extensions | v1beta1 | SupplementalGroupsStrategyOptions



SupplementalGroupsStrategyOptions defines the strategy type and options used to create the strategy.

<aside class="notice">
Appears In  <a href="#podsecuritypolicyspec-v1beta1">PodSecurityPolicySpec</a> </aside>

Field        | Description
------------ | -----------
ranges <br /> *[IDRange](#idrange-v1beta1) array*  | Ranges are the allowed ranges of supplemental groups.  If you would like to force a single supplemental group then supply a single range with the same start and end.
rule <br /> *string*  | Rule is the strategy that will dictate what supplemental groups is used in the SecurityContext.

