# Version 0.2.0 Change Log
The version 0.2.0 reflects changes in the Lachesis API v.0.7.0-rc1 and changes in SFC contract v.2.0.2-rc1 deployed into the Fantom Opera main net on 8/2020.

The API implementation follows changes in registering, processing and reporting of stake delegations. The new Lachesis and SFC version implements so called fluid staking model. Each delegator has ability to delegate multiple times from a single address. A delegation can be locked for a time period chosen by the delegator to gain better rewards for the delegation locked.

The original 1:1 relationship between network account address and delegation no longer applies and has been changed to 1:n relationship. All API endpoints working with the delegation reflect this change.

## GraphQL Schema Changes List

- `Delegator` type has been renamed to `Delegation`.
- `DelegatorList` type has been renamed to `DelegationList`.
- Attribute `delegator` of type `DelegationListEdge` has been renamed to `delegation`.
- Attribute `delegation` of type `Account` has been renamed to `delegations`.
- Type of attribute `delegations` of type `Account` has been changed to `DelegationList`.
- `DelegationList` control options have been added to attribute `delegation` of type `Account`. The attribute structure is now `delegations(cursor:Cursor, count:Int = 25): DelegationList!`.
- New mandatory attribute `staker` of type `Long` has been added to the type `PendingRewards`.
- Type of attribute `delegations` of type `Staker` has been changed to `DelegationList` to reflect the delegation list type name change.
- Query attribute `delegationsOf` type has been changed to `DelegationList` to reflect delegation list type name change.
- Query attribute `delegation` parameters now include mandatory option `staker` of type `Long` to specify unique delegation in 1:n relationship.
- Query attribute `delegation` type has been changed to `Delegation` to reflect a single delegation type name change.
- New query attribute `delegationsByAddress` of type `DelegationList` has been added to expose list of all delegations of a single delegator address. The attribute call definition is `delegationsByAddress(address:Address!, cursor: Cursor, count: Int = 25): DelegationList!`.
