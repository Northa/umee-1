package e2e

import (
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
)

var (
	ATOM_SYMBOL     = "ATOM"
	ATOM_BASE_DENOM = "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2"
	ATOM_EXPONENT   = 6
)

func createMnemonic() (string, error) {
	entropySeed, err := bip39.NewEntropy(256)
	if err != nil {
		return "", err
	}

	mnemonic, err := bip39.NewMnemonic(entropySeed)
	if err != nil {
		return "", err
	}

	return mnemonic, nil
}

func createMemoryKey() (mnemonic string, info *keyring.Record, err error) {
	mnemonic, err = createMnemonic()
	if err != nil {
		return "", nil, err
	}

	account, err := createMemoryKeyFromMnemonic(mnemonic)
	if err != nil {
		return "", nil, err
	}

	return mnemonic, account, nil
}

func createMemoryKeyFromMnemonic(mnemonic string) (*keyring.Record, error) {
	kb, err := keyring.New("testnet", keyring.BackendMemory, "", nil, cdc)
	if err != nil {
		return nil, err
	}

	keyringAlgos, _ := kb.SupportedAlgorithms()
	algo, err := keyring.NewSigningAlgoFromString(string(hd.Secp256k1Type), keyringAlgos)
	if err != nil {
		return nil, err
	}

	account, err := kb.NewAccount("", mnemonic, "", sdk.FullFundraiserPath, algo)
	if err != nil {
		return nil, err
	}

	return account, nil
}
