package cloudwalletclass



type EthereumWalletAcc struct{
	Email string `json:"Email"`
	Password string `json:"Password"`
	EthAddress string `json:"EthAddress"`
	Terms bool `json:"Terms"`
	Allowed bool `json:"Allowed"`
	PrvteKey string `json:"PrvteKey"`
}



