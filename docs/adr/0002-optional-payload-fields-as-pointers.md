# 2. Optional payload fields as pointers

Date: 2021-02-16

## Status

Accepted

## Context

Kentik HTTP API accepts and returns some payload's fields optionally, it means:
- the API accepts post requests with some fields missing and so they are set to their zero/null values
- the API accepts put requests with only selected fields and only the specified fields are being updated
- the API responds with some fields set to null or some fields missing
GO has no concept of Optional and this shortcoming needs to be addressed.

## Decision

Declare optional fields in payload structs as pointers.

## Consequences

The field-selective HTTP API interface can be fully exposed to the library user.
