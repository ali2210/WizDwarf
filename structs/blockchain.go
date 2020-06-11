package structs

import(
	"math/big"
)



// Hold Block properties
type Block struct{
	Balance *big.Int 
	TxSen string  // whom to send ...
	TxRec string  // whom to receive....
	FeesCharges string
	Nonce string
	GasPrice string
}