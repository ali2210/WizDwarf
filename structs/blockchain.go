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
	Nonce uint64
	GasPrice *big.Int
	GasLimit uint64
}

