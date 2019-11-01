---
title: Problem Details For CoAP APIs
abbrev: CoAP Problem
docname: draft-fossati-core-coap-problem-latest
category: STD

ipr: trust200902
area: ART
workgroup: CoRE Working Group
keyword: CoAP, API, Problem Details

stand_alone: yes
pi: [toc, sortrefs, symrefs]

author:
 -
    ins: "T. Fossati"
    name: "Thomas Fossati"
    organization: "ARM"
    email: thomas.fossati@arm.com

normative:
  RFC7049: cbor
  RFC7252: coap
  RFC8610: cddl

informative:
  RFC7807:

--- abstract

This document defines a "problem detail" as a way to carry machine-readable
details of errors in a CoAP response to avoid the need to define new error
response formats for CoAP APIs.

--- middle

# Introduction

CoAP {{coap}} response codes are sometimes not sufficient to convey enough
information about an error to be helpful.

This specification defines simple and extensible CBOR {{cbor}} format to suit
this purpose.  It is designed to be reused by CoAP APIs, which can identify
distinct "problem types" specific to their needs.

Thus, API clients can be informed of both the high-level error class (using
the response code) and the finer-grained details of the problem (using this
format).

The format presented is largely inspired by the Problem Details for HTTP APIs
defined in {{RFC7807}}.

## Requirements Language

{::boilerplate bcp14}


# Security Considerations

TODO Security


# IANA Considerations

This document has no IANA actions.



--- back

# Acknowledgments
{:numbered="false"}

TODO acknowledge.
