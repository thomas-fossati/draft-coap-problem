#!/bin/bash

set -o pipefail
set -eu

# expects:
#   t.Logf("HexEncoded # %s\n", got)
go test -v | 
  grep 'HexEncoded' | 
  cut -f2 -d'#' | 
  while read e ; \
  do \
    echo ${e} | xxd -r -p | cddl ../coap-problem-details.cddl validate - ;
    if [ $? -ne 0 ]; \
    then \
      printf "[!] CDDL validation failed for %s\n" ${e} ; \
    else \
      echo ${e} ; \
      echo ${e} | xxd -r -p | cbor2diag.rb ; \
    fi
  done

# vim: set ts=2 et:
