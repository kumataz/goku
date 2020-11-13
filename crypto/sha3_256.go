/*
use by call linux cmd to sha3-256 hash ，version: FIPS standard keccake-256
linux cmd tool need install libc6-dev
dell data type： []uint8 -> []uint8
*/
package crypto

import (
	"os/exec"
	"fmt"
	"io/ioutil"
	"strings"
	"encoding/hex"
)

/*  simple data

var INPUT = [55]byte{
    0x52, 0xe7, 0xa3, 0x05,
    0x90, 0x2f, 0xa9, 0xfb,
    0xf6, 0x24, 0x20, 0xa7,
    0x84, 0x96, 0x87, 0x2f,
    0x57, 0xe7, 0x2b, 0xe6,
    0x32, 0x15, 0x99, 0x14,
    0x89, 0x20, 0x9a, 0x84,
    0xd4, 0x50, 0x5b, 0x1a,
    0x00, 0x62, 0x82, 0x5b,
    0x00, 0x00, 0x00, 0x00,
    0x00, 0x00, 0x00, 0x00,
    0x00, 0x00, 0x00, 0x00,
    0x00, 0x00, 0x04, 0x5b,
	0x00, 0x05, 0xe9,
}

// "abc"
var TEST = [3]byte{
	0x61, 0x62, 0x63,
}
*/

func Sha3sum256(data []byte) []byte{
	input := strings.Join([]string{"echo", hex.EncodeToString(data), "| xxd -r -p | sha3sum -a 256"}," ")
	cmd := exec.Command("/bin/bash", "-c", input)
	// Create get command output pipeline
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
	}
	
	// execute cmd
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
	}
	
	// read all ouput
	result, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
	}
	
	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
	}
	
	// output
	checkhash, _ := hex.DecodeString(string(result[:64]))
	return checkhash

}

/*
func main(){
	sha3_hash := crypto.Sha3sum256(INPUT[:])
	fmt.Printf("SHA3-256:   %+02x\n", sha3_hash)
}
*/