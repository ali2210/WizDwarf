package structs

import(
	"math/big"
)



// Hold Block properties
type Block struct{
	Balance *big.Int 
	TxSen string  // whom to send ...
	TxRec string  // our customer
	FeesCharges string
	Nonce int32
	GasPrice int32
}

