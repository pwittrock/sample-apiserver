## FSGroupStrategyOptions v1beta1 extensions

Group        | Version     | Kind
------------ | ---------- | -----------
extensions | v1beta1 | FSGroupStrategyOptions



FSGroupStrategyOptions defines the strategy type and options used to create the strategy.

<aside class="notice">
Appears In  <a href="#podsecuritypolicyspec-v1beta1">PodSecurityPolicySpec</a> </aside>

Field        | Description
------------ | -----------
ranges <br /> *[IDRange](#idrange-v1beta1) array*  | Ranges are the allowed ranges of fs groups.  If you would like to force a single fs group then supply a single range with the same start and end.
rule <br /> *string*  | Rule is the strategy that will dictate what FSGroup is used in the SecurityContext.

