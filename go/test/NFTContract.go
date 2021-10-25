package test

import (
	"regexp"
	"testing"

	emulator "github.com/onflow/flow-emulator"
	"github.com/onflow/flow-go-sdk"
	sdk "github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-go-sdk/templates"
	sdktemplates "github.com/onflow/flow-go-sdk/templates"
	"github.com/onflow/flow-go-sdk/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ft_contracts "github.com/onflow/flow-ft/lib/go/contracts"
	nft_contracts "github.com/onflow/flow-nft/lib/go/contracts"
)

const (
	nowwhereRootPath           = "../.."
	NFTContractPath            = nowwhereRootPath + "/contracts/NFTContract.cdc"
	NowwhereSetupAccountPath   = nowwhereRootPath + "/transactions/kibble/setup_account.cdc"
	nowwhereTransferTokensPath = nowwhereRootPath + "/transactions/kibble/transfer_tokens.cdc"
	nowwhereMintTokensPath     = nowwhereRootPath + "/transactions/kibble/mint_tokens.cdc"
	nowwhereBurnTokensPath     = nowwhereRootPath + "/transactions/kibble/burn_tokens.cdc"
	nowwhereGetBalancePath     = nowwhereRootPath + "/scripts/kibble/get_balance.cdc"
	nowwhereGetSupplyPath      = nowwhereRootPath + "/scripts/getSupply.cdc"
	nowwhereGetCollectionPath  = nowwhereRootPath + "/scripts/getCollection.cdc"
)

func NowwhereDeployContracts(b *emulator.Blockchain, t *testing.T) (flow.Address, flow.Address, crypto.Signer) {
	accountKeys := test.AccountKeyGenerator()

	nftCode := loadNonFungibleToken()
	nftAddr, err := b.CreateAccount(
		nil,
		[]sdktemplates.Contract{
			{
				Name:   "NonFungibleToken",
				Source: string(nftCode),
			},
		},
	)
	require.NoError(t, err)

	////////////////
	_, err = b.CommitBlock()
	assert.NoError(t, err)

	nowwhereAccountKey, nowwhereSigner := accountKeys.NewWithSigner()
	NFTContractCode := loadNowwhereNFT(nftAddr.String())

	nowwhereAddr, err := b.CreateAccount(
		[]*flow.AccountKey{nowwhereAccountKey},
		[]templates.Contract{templates.Contract{
			Name:   "NFTContract",
			Source: string(NFTContractCode),
		}},
	)
	assert.NoError(t, err)

	_, err = b.CommitBlock()
	assert.NoError(t, err)

	// Simplify testing by having the contract address also be our initial Vault.
	//	NowwhereSetupAccount(t, b, nowwhereAddr, nowwhereSigner, fungibleAddr, nowwhereAddr)

	return nftAddr, nowwhereAddr, nowwhereSigner
}

func NowwhereSetupAccount(t *testing.T, b *emulator.Blockchain, userAddress sdk.Address, userSigner crypto.Signer, fungibleAddr sdk.Address, nowwhereAddr sdk.Address) {
	tx := flow.NewTransaction().
		SetScript(nowwhereGenerateSetupKibbleAccountTransaction(fungibleAddr, nowwhereAddr)).
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().Index, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(userAddress)

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, userAddress},
		[]crypto.Signer{b.ServiceKey().Signer(), userSigner},
		false,
	)
}

func nowwhereReplaceAddressPlaceholders(code string, nonfungibleAddress, nftContractAddress string) []byte {
	return []byte(replaceImports(
		code,
		map[string]*regexp.Regexp{
			nonfungibleAddress: nftAddressPlaceholder,
			nftContractAddress: NFTContractAddressPlaceHolder,
		},
	))
}

func loadFungibleToken() []byte {
	return ft_contracts.FungibleToken()
}

func loadNowwhereNFT(nftAddr string) []byte {
	return []byte(replaceImports(
		string(readFile(NFTContractPath)),
		map[string]*regexp.Regexp{
			nftAddr: nftAddressPlaceholder,
		},
	))
}

func loadNFT(fungibleAddr flow.Address) []byte {
	return []byte(replaceImports(
		string(readFile(NFTContractPath)),
		map[string]*regexp.Regexp{
			fungibleAddr.String(): ftAddressPlaceholder,
		},
	))
}

func NowwhereGenerateGetSupplyScript(fungibleAddr, nowwhereAddr flow.Address) []byte {
	return nowwhereReplaceAddressPlaceholders(
		string(readFile(nowwhereGetSupplyPath)),
		fungibleAddr.String(),
		nowwhereAddr.String(),
	)
}

func NowwhereGenerateGetCollectionScript(fungibleAddr, nowwhereAddr flow.Address) []byte {
	return nowwhereReplaceAddressPlaceholders(
		string(readFile(nowwhereGetCollectionPath)),
		fungibleAddr.String(),
		nowwhereAddr.String(),
	)
}

func nowwhereGenerateSetupKibbleAccountTransaction(fungibleAddr, nowwhereAddr flow.Address) []byte {
	return nowwhereReplaceAddressPlaceholders(
		string(readFile(NowwhereSetupAccountPath)),
		fungibleAddr.String(),
		nowwhereAddr.String(),
	)
}

func loadNonFungibleToken() []byte {
	return nft_contracts.NonFungibleToken()
}
