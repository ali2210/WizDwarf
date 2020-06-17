package cloudwalletclass



/*import(
	 "crypto/ecdsa"
	"fmt"
)*/


// var WalletPrivateKey  string = ""

type EthereumWalletAcc struct{
	Email string `json:"Email"`
	Password string `json:"Password"`
	EthAddress string `json:"EthAddress"`
	Terms bool `json:"Terms"`
	// PubKey string `json:"PubKey"`
	PrvteKey string `json:"PrvteKey"`
}



// func (a *EthereumWalletAcc)SetPrivateKey(){
// 	WalletPrivateKey = (*a).PrvteKey 
// 	fmt.Println("WalletPrivateKey:", WalletPrivateKey)
// }

// func(*EthereumWalletAcc)GetPrivateKey()(string){

// 	if WalletPrivateKey == ""{
// 	return WalletPrivateKey
// 	}
//  return WalletPrivateKey
// }
