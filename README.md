## Commands

## Start Flow

### Creating the contract and minting a token
flow project start-emulator

flow project deploy

flow keys generate

## Create Template argument is max supply
flow transactions send transactions/CreateTemplateData.cdc --arg UInt32:2  --network testnet --signer testnet-account


## Mint NFT argument template ID
flow transactions send transactions/mint.cdc --arg UInt64:2  --network testnet --signer testnet-account

## Command to lock template argument template ID

flow transactions send transactions/locktemplate.cdc --arg UInt64:1  --network testnet --signer testnet-account

## Command to unlock template argument template ID

flow transactions send transactions/unlocktemplate.cdc --arg UInt64:1  --network testnet --signer testnet-account# flowSmartcontracts
# flowSmartcontracts
