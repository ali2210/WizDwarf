package structs


import(
	"crypto/ecdsa"
)

type SignedKey struct {
	Reader string
	Signed string
	Tx *ecdsa.PrivateKey
}