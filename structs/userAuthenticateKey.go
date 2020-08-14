package structs


import(
	"crypto/ecdsa"
)

type SignedKey struct {
	reader string
	signed string
	tx *ecdsa.PrivateKey
}