# Abstract

This document defines a "problem detail" as a way to carry machine-readable
details of errors in a CoAP response to avoid the need to define new error
response formats for CoAP APIs.


# CDDL

The coap-problem-details-extension CDDL socket is used to add new information
structures to the coap-problem-details root map.

~~~
coap-problem-details = {
  type => uint,
  ? title => text,
  ? response-code => bstr .size 1,
  ? detail => text,
  ? instance => uri,
  * $$coap-problem-details-extension,
}

type = 0
title = 1
response-code = 2
detail = 3
instance = 4
~~~

# IANA

## New Content-Format

This document requests that IANA registers the following Content-Format to the
"CoAP Content-Formats" subregistry, within the "Constrained RESTful
Environments (CoRE) Parameters" registry, from the Expert Review space
(0..255):
~~~
       +-------------------------------+----------+------+-----------+
       | Media Type                    | Encoding | ID   | Reference |
       +-------------------------------+----------+------+-----------+
       | application/coap-problem+cbor | --       | TBD1 | RFCthis   |
       +-------------------------------+----------+------+-----------+
~~~

## New Registries

This document requests that IANA create the following new registries:

o  CoAP Problem Details (Section ...)
o  CoAP Problem Types (Section ...)

### CoAP Problem Details Registry

The "CoAP Problem Details" registry keeps track of the allocation of the
integer values used as index values in the coap-problem-details CBOR map.

Future registrations for this registry are to be made based on [RFC8126] as
follows:
~~~
	      +------------------+-------------------------+
              | Range            | Registration Procedures |
              +------------------+-------------------------+
              | 0-N              | Standards Action        |
              |                  |                         |
              | N+1-4294967295   | Specification Required  |
              +------------------+-------------------------+

                Table 1:  CoAP Problem Details Procedures
~~~

All negative values are reserved for Private Use.

Initial registrations for the "CoAP Problem Details" registry are provided
below.  Assignments consist of an integer index value, the item name, and a
reference to the defining specification.
~~~
       +---------------+---------------------------+---------------+
       | Index         | Item Name                 | Specification |
       +---------------+---------------------------+---------------+
       | 0             | type                      | RFC-THIS      |
       |               |                           |               |
       | 1             | title                     | RFC-THIS      |
       |               |                           |               |
       | 2             | response-code             | RFC-THIS      |
       |               |                           |               |
       | 3             | detail                    | RFC-THIS      |
       |               |                           |               |
       | 4             | instance                  | RFC-THIS      |
       +---------------+---------------------------+---------------+

             Table 2: CoAP Problem Details Inital Registrations
~~~

### CoAP Problem Types Registry

The "CoAP Problem Details" registry keeps track of the problem type values.

Future registrations for this registry are to be made based on [RFC8126] as
follows:
~~~
	      +------------------+-------------------------+
              | Range            | Registration Procedures |
              +------------------+-------------------------+
              | 0-M              | Standards Action        |
              |                  |                         |
              | M+1-4294967295   | Specification Required  |
              +------------------+-------------------------+

                Table 3:  CoAP Problem Details Proceedures
~~~

This specification reserves the use of one value as a problem type:

The 0 value, when used as a problem type, indicates that the problem has no
additional semantics beyond that of the CoAP response code.

The initial registration for the "CoAP Problem Types" registry is provided
below.  Assignments consist of an integer index value, the item name, and a
reference to the defining specification.
~~~
  +---------------+-------------------------------+---------------+
  | Value         | Description                   | Specification |
  +---------------+-------------------------------+---------------+
  | 0             | The problem has no additional | RFC-THIS      |
  |               | semantics beyond that of the  |               |
  |               | CoAP response code            |               |
  +---------------+-------------------------------+---------------+
 
             Table 4: CoAP Problem Types Inital Registration
~~~
