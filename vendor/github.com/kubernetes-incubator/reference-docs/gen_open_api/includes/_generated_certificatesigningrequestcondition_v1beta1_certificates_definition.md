## CertificateSigningRequestCondition v1beta1 certificates

Group        | Version     | Kind
------------ | ---------- | -----------
certificates | v1beta1 | CertificateSigningRequestCondition





<aside class="notice">
Appears In  <a href="#certificatesigningrequeststatus-v1beta1">CertificateSigningRequestStatus</a> </aside>

Field        | Description
------------ | -----------
lastUpdateTime <br /> *[Time](#time-v1)*  | timestamp for the last update to this condition
message <br /> *string*  | human readable message with details about the request state
reason <br /> *string*  | brief reason for the request state
type <br /> *string*  | request approval state, currently Approved or Denied.

