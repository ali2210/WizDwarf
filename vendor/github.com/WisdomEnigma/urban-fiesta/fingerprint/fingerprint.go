package fingerprint

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

type Image_Print struct{}

func (*Image_Print) CalculateHashColor(r, g, b, k uint32, _sizeof int64) string {

	hash_chain := make([][4]string, _sizeof)

	var encode_chain [32]byte
	for i := 0; i < len(hash_chain); i++ {

		hash_chain[i][0] = strconv.FormatUint(uint64(r), 10)
		hash_chain[i][1] = strconv.FormatUint(uint64(g), 10)
		hash_chain[i][2] = strconv.FormatUint(uint64(b), 10)
		hash_chain[i][3] = strconv.FormatUint(uint64(k), 10)
	}

	for i := range hash_chain {
		for j := 0; j < 4; j++ {
			encode_chain = sha256.Sum256([]byte(hash_chain[i][j]))
		}
	}

	return hex.EncodeToString(encode_chain[:])

}
