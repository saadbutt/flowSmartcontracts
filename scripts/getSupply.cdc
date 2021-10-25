import NFTContract from "../contracts/NFTContract.cdc"

// This script returns the total amount of NFTs currently in existence.

pub fun main(): UInt64 {

    let supply = NFTContract.totalSupply

    log(supply)

    return supply
}
