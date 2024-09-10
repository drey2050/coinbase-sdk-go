# Coinbase Go SDK Changelog

## [0.0.6] - 2024-08-27

- Adds headers to requests with SDK version & language

## [0.0.2] - 2024-08-27

### Fixed

- Fixed a bug where we weren't handling the api returned validators properly resulting in an index out of range error.

## [0.0.1] - 2024-08-23

### Added

Initial release of the Coinbase Go SDK

- Support for Staking on External Wallets
    - Full support for Shared ETH Staking
    - Partial support for Dedicated ETH Staking
      - Only stake is supported, unstake will be coming soon
    - On networks `holesky` and `mainnet`
- Support for getting stakeable balances on External Wallets
- Support for Listing Rewards on External Wallets
