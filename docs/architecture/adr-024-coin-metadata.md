# ADR 024: Coin Metadata

## Changelog

- 05/19/2020: Initial draft

## Status

Proposed

## Context

Assets in the Cosmos SDK are represented via a `Coins` type that consists of an `amount` and a `denom`,
where the `amount` can be any arbitrarily large or small value. In addition, the Cosmos SDK uses an
account-based model where there are two types of primary accounts -- basic accounts and module accounts.
All account types have a set of balances that are composed of `Coins`. The `x/bank` module keeps
track of all balances for all accounts and also keeps track of the total supply of balances in an
application.

With regards to a balance `amount`, the Cosmos SDK assumes a static and fixed unit of denomination,
regardless of the denomination itself. In other words, clients and apps built atop a Cosmos-SDK-based
chain may chose to define and use arbitrary units of denomination to provide a richer UX, however, by
the time a tx or operation reaches the Cosmos SDK state machine, the `amount` is treated as a single
unit. For example, for the Cosmos Hub (Gaia), clients assume 1 ATOM = 10^6 uatom, and so all txs and
operations in the Cosmos SDK work off of units of 10^6.

This clearly provides a poor and limited UX especially as interoperability of networks increases and
as a result the total amount of asset types increases. We propose to have `x/bank` additionally keep
track of metadata per `denom` in order to help clients, wallet providers, and explorers improve their
UX and remove the requirement for making any assumptions on the unit of denomination.

## Decision

The `x/bank` module will updated to store and index metadata by `denom`, specifically the "base" or
smallest unit -- the unit the Cosmos SDK state-machine works with. The metadata should ideally be
extensible and serve the needs of many external consumers and even other modules (e.g. `x/ibc` or `x/nft`).
At a minimum, the metadata type should include all units and a potential list of denomination aliases.
As a result, we can define the type as follows:

```protobuf
message DenomUnit {
  string denom    = 1;
  uint32 exponent = 2;  
}

message Metadata {
  string description = 1;
  string base = 2;
  repeated DenomUnit denom_units = 3;
  repeated string aliases = 4;
}
```

As an example, the ATOM's metadata can be defined as follows:

```json
{
  "description": "The native staking token of the Cosmos Hub.",
  "base": "uatom",
  "denom_units": [
    {
      "denom": "matom",
      "exponent": 3,
    },
    {
      "denom": "atom",
      "exponent": 6,
    }
  ],
  "aliases": [
    "UATOM",
    "microatom"
  ]
}
```

> Given the above metadata, a client may infer that 4.3atom = 4.3 * (10^6) = 4,300,000uatom.

A client should be able to query for metadata by denom both via the CLI and REST interfaces. In
addition, we will add handlers to these interfaces to convert from any unit to another given unit,
as the base framework for this already exists in the Cosmos SDK.

## Future Work

In order for clients to avoid having to convert assets to the base denomination -- either manually or
via an endpoint, we may consider supporting automatic conversion of a given unit input.

## Consequences

### Positive

- Provides clients, wallet providers and block explorers with additional data on
  asset denomination to improve UX and remove any need to make assumptions on
  denomination units.

### Negative

- A small amount of required additional storage in the `x/bank` module. The amount
  of additional storage should be minimal as the amount of total assets should not
  be large.

### Neutral

## References