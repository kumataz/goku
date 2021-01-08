package crypto

import (
	"golang.org/x/crypto/sha3"
)


// New256 creates a new SHA3-256 hash.
func Sha3_256(data ...[]byte) []byte {
	hash := sha3.New256()
	for _, b := range data {
		hash.Write(b)
	}
	return hash.Sum(nil)
}

// Keccak256 calculates and returns the Keccak256 hash of the input data.
// NewLegacyKeccak256 creates a new Keccak-256 hash.
func Keccak256(data ...[]byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	for _, b := range data {
		hash.Write(b)
	}
	return hash.Sum(nil)
}


// // "abc"
// var TEST = [3]byte{
// 	0x61, 0x62, 0x63,
// }

// func main(){
// 	keccak_hash := crypto.Keccak256(TEST[:])
// 	fmt.Printf("Keccak-256: %+02x\n", keccak_hash)
//
//     sha3_hash := crypto.Sha3_256(TEST[:])
//     fmt.Printf("  SHA3-256: %+02x\n", sha3_hash)
// }
