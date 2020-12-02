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



// // "abc"
// var TEST = [3]byte{
// 	0x61, 0x62, 0x63,
// }


func Sha3_256_LinuxCMD(data []byte) []byte{
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


// func main(){
// 	sha3_hash := crypto.Sha3_256_LinuxCMD(TEST[:])
// 	fmt.Printf("SHA3-256:   %+02x\n", sha3_hash)
// }
