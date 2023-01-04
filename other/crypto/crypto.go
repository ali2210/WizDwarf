package crypto

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"log"
	"reflect"

	linkcid "github.com/ipfs/go-cid"
	multihash "github.com/multiformats/go-multihash"
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

func FilePrints(prints ...string) (map[string]string, linkcid.Cid) {

	nullify := map[string]string{}
	hash_data := sha256.Sum256([]byte(prints[0]))

	// this function also high order function which convert hash of bytes into encoded string
	// fog(c) = f(g(x)) mathematical notion of high order function
	// encoded string in hex format which mus be decode as string in hex format
	decoder, err := hex.DecodeString(hex.EncodeToString(hash_data[:]))
	if err != nil {

		log.Fatalln("Error:", err)
		return nullify, linkcid.Cid{}
	}

	// EncodeName function takes decoder which is already in hex format and then apply x11 crypto algorithm.
	// with x11 breakable signature into unbreakable
	encodetype, err := multihash.EncodeName(decoder, "x11")
	if err != nil {

		log.Fatalln("Error:", err)
		return nullify, linkcid.Cid{}
	}

	// This is an higher order function encodeName value as encode string in hex format.
	// multihash hex string apply on encoded string in hex format.
	encodex11, err := multihash.FromHexString(hex.EncodeToString(encodetype))
	if err != nil {

		log.Fatalln("Error:", err)
		return nullify, linkcid.Cid{}
	}

	// generate new cid.. The specification of this function require two parameters (codeType & other one is hash algorithm)
	// merkel tree (dag) data serilaization (protocol buffer [https://en.wikipedia.org/wiki/Protocol_Buffers])
	// & hash algorithm

	cid := linkcid.NewCidV1(linkcid.DagProtobuf, encodex11)

	// check whether cid version is 0. For this application cid version must be 1
	if version := cid.Version(); version != 1 {

		log.Fatalln("Error:", err)
		return nullify, linkcid.Cid{}
	}

	// create protocol buffer map object
	// map key which we had calculated before cid because key must be in string
	// value should be cdr link

	cdr := make(map[string]string, 1)

	cdr[cid.String()] = cid.Hash().B58String()

	return cdr, cid
}
