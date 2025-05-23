package utils

import (
	"fmt"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/remiehneppo/go-binance-api/types"
)

func CreateHDWalletByMnemonic(mnemonic string, index int) (privateKey, address string, err error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		fmt.Println("Error creating wallet from mnemonic:", err)
		return
	}

	path := hdwallet.MustParseDerivationPath(fmt.Sprintf("%s%d", types.HD_BASE_PATH, index))
	account, err := wallet.Derive(path, false)
	if err != nil {
		fmt.Println("Error deriving account:", err)
		return
	}
	privateKey, err = wallet.PrivateKeyHex(account)
	if err != nil {
		fmt.Println("Error getting private key:", err)
		return
	}
	address, err = wallet.AddressHex(account)
	if err != nil {
		fmt.Println("Error getting address:", err)
		return
	}
	return privateKey, address, nil
}
