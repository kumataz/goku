
package main

import(
    "fmt"
    "goku/internal/crypto"
    // . "goku/basics"
)

// "abc"
var TEST = [3]byte{
	0x61, 0x62, 0x63,
}


func main(){
	keccak_hash := crypto.Keccak256(TEST[:])
	fmt.Printf("Keccak-256: %+02x\n", keccak_hash)

    sha3_hash := crypto.Sha3_256(TEST[:])
    fmt.Printf("  SHA3-256: %+02x\n", sha3_hash)
}
