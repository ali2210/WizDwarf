package structs

import(
	// "crypto/ecdsa"
	"fmt"
)


var WalletPrivateKey  string = ""

type Acc struct{
	Email string
	Password string
	Terms bool
	EthAddress string
	PubKey string
	PrvteKey string 
}


func (a *Acc)SetPrivateKey(){
	WalletPrivateKey = (*a).PrvteKey 
	fmt.Println("WalletPrivateKey:", WalletPrivateKey)
}

func(*Acc)GetPrivateKey()(string){

	if WalletPrivateKey == ""{
	return WalletPrivateKey
	}
 return WalletPrivateKey
}
