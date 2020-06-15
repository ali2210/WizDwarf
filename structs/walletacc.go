package structs

import(
	"crypto/ecdsa"
	"fmt"
)


var WalletPrivateKey  *ecdsa.PrivateKey = nil

type Acc struct{
	Email string
	Password string
	Terms bool
	EthAddress string
	PubKey string
	PrvteKey *ecdsa.PrivateKey 
}


func (a *Acc)SetPrivateKey(){
	WalletPrivateKey = (*a).PrvteKey 
	fmt.Println("WalletPrivateKey:", WalletPrivateKey)
}

func(*Acc)GetPrivateKey()(*ecdsa.PrivateKey){

	if WalletPrivateKey == nil{
	return nil
	}
 return WalletPrivateKey
}