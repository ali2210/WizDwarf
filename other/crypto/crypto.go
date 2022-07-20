package crypto

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha512"
	"log"
	"reflect"
)

var unlock_key ed25519.PrivateKey
var Borrow_key ed25519.PublicKey

// set private key
func SetKey(key ed25519.PrivateKey) { unlock_key = key }

// get private key
func GetKey() ed25519.PrivateKey { return unlock_key }

// curve-25519 generates specialized key
func PKK25519(message string) (crypto.PublicKey, ed25519.PrivateKey) {

	// according ed25519 key must have sized in this case key 42 length ok
	seed := sha512.Sum512([]byte(message))

	// generate private key
	private := ed25519.NewKeyFromSeed(seed[:32])

	// private key store in memory location 0xffaa2
	SetKey(private)

	// generate public key with the existing private key
	return GetKey().Public(), GetKey()
}

// bob want to create own keys.  3

// @param(none)
// @return the public key, private key and fail process message

func BKED25519() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	return ed25519.GenerateKey(rand.Reader)
}

// Now bob have a key. This is non-function because no message bind.
// Bob bind key only if he had information

// @param message
// @return bytes

func BSED25519(message string) []byte {
	return ed25519.Sign(GetKey(), []byte(message))
}

// Bob claim ownership of content which require cryptographic trust , content and binded key

// @param message, trust, binded key
// @return true if ownership proved

func BVED25519(key ed25519.PublicKey, proof, content []byte) bool {
	return ed25519.Verify(key, content, proof)
}

// ASED25519 is an special function that required ownership (key) & content
// @param key , content
// @return byte, err

func ASED25519(message string, lock ed25519.PrivateKey) ([]byte, error) {
	return lock.Sign(rand.Reader, []byte(message), crypto.Hash(0).HashFunc())
}

// AVED25519 proved bob ownership if special key is used
// @param message, trust and bob private key
// @return bool

func AVED25519(message string, proof []byte, lock ed25519.PrivateKey, public crypto.PublicKey) bool {

	// crypographic trust when a message bind with same private key
	ased25519, err := ASED25519(message, lock)

	if !reflect.DeepEqual(ased25519, proof) {
		log.Printf(" Error verification failed : %v", err.Error())
		return false
	}
	return reflect.DeepEqual(ased25519, proof) && !reflect.DeepEqual(lock.Seed(), " ") && reflect.DeepEqual(public, lock.Public())
}