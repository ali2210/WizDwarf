package structs

// "crypto/ecdsa"

var WalletPrivateKey string = ""

type Acc struct {
	Email      string
	Password   string
	Terms      bool
	EthAddress string
	PubKey     string
	PrvteKey   string
	Allowed    bool
}

func (a *Acc) SetPrivateKey() {

	WalletPrivateKey = (*a).PrvteKey

}

func (*Acc) GetPrivateKey() string {

	if WalletPrivateKey == "" {
		return WalletPrivateKey
	}
	return WalletPrivateKey
}
