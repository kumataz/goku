// 调用Linux命令行返回输出结果
package pkgtools

import (
	"os/exec"
	"fmt"
	"io/ioutil"
)


func main() {
	input := "ls -l"
	cmd := exec.Command("/bin/bash", "-c", input)
	// Create get command output pipeline
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
	return
	}
	
	// execute cmd
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
	return
	}
	
	// read all ouput
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
	return
	}
	
	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
	return
	}
	
	// output
	fmt.Printf("%s\n", bytes)
}
