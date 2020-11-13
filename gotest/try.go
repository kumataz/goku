package gotest

import (
	"time"
	"fmt"
	"github.com/terrancewong/serial"
	"encoding/hex"
)

// test1
func Sum(numbers []int) int {  
	 sum := 0
	 // 这个 bug 是故意的
	 for _,n := range numbers {
		 sum += n
	 }
	 return sum
}


// test2
var cfg_port = "/dev/ttyUSB0"
var work_port = "/dev/ttyUSB1"

func openSer(s string) *serial.Port {
    conf_ser := &serial.Config{Name: s, Baud: 115200, StopBits: serial.Stop2, ReadTimeout: 4*time.Millisecond}
    ser, err := serial.OpenPort(conf_ser)
    if err != nil {
        //log.Fatal(err)
        panic(err)
    }
    return ser
}


// TODO write to port
func cfgWriter(ser *serial.Port, setData string){
    fmt.Println("Send argv at", time.Now().Format("2006-01-02 15:04:05"), cfg_port)
    s, _ := hex.DecodeString(setData)
	p := []byte(s)
	_, err := ser.Write(p) 
	
	if err != nil{
		fmt.Println("error: ", err)
	}else{
		fmt.Printf( ">> %x \n", p)
	}
	
	ser.Flush()
	return
}
