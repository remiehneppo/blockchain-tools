package types

type RootWallet struct {
	Mnemonic   string `json:"mnemonic"`
	PrivateKey string `json:"private_key"`
	Address    string `json:"address"`
}
