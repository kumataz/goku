package embedded

import (
	"time"
	"fmt"
	"github.com/terrancewong/serial"
	"encoding/hex"
)

var cfg_port = "/dev/ttyUSB0"
var work_port = "/dev/ttyUSB1"

func OpenSer(s string) *serial.Port {
    conf_ser := &serial.Config{Name: s, Baud: 115200, StopBits: serial.Stop2, ReadTimeout: 4*time.Millisecond}
    ser, err := serial.OpenPort(conf_ser)
    if err != nil {
        //log.Fatal(err)
        panic(err)
    }
    return ser
}


// TODO write to port
func CfgWriter(ser *serial.Port, setData string){
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


// TODO read from port
func CfgReader(ser *serial.Port){
	time.Sleep(2 * time.Millisecond)	
	buf := make([]byte, 16)
//	n := 1
//	// clear rx fifo garbage
//	for ; n > 0 ;n, _ = ser.Read(buf){
//	}
    for {
        if n, err := ser.Read(buf) ; err != nil {
        	
        } else {
        	if (n == 16) {
		        fmt.Printf( "<< %x\n", buf[:n])
        	}
        }
    }
}


func SwitchPort(port string){
	ser :=  OpenSer(port)
	defer ser.Close()
	
	// write to port
	s, _ := hex.DecodeString("aa5aff00a0ff0000005cfffe90ffff00a0ff00010064000090ffff01")
	p := []byte(s)
	_, err := ser.Write(p) 
	if err != nil{
		fmt.Println("error: ", err)
	}
	ser.Flush()
	
	// read from port	
	buf := make([]byte, 8)
    for {
        if n, rerr := ser.Read(buf) ; rerr != nil {
        	
        } else {
        	if (n == 8  && "90" == hex.EncodeToString(buf[:1])){	
				fmt.Printf( "cfg_port:  %s \n", cfg_port)
				fmt.Printf("work_port: %s \n", work_port)
				return
        	}else{
        		cfg_port = "/dev/ttyUSB1"
				work_port = "/dev/ttyUSB0"
				fmt.Printf( "cfg_port:  %s \n", cfg_port)
				fmt.Printf("work_port: %s \n", work_port)
				return
        	}
        }
    }
}

	
//func main(){
//	SwitchPort(cfg_port)
//	cfg_ser := OpenSer(cfg_port)  
//	defer cfg_ser.Close()
//	setData := "aa5aff00a0ff0000005cfffe90ffff00a0ff00010064000090ffff01"
//	go CfgWriter(cfg_ser, setData)
//	go CfgReader(cfg_ser)
//	
//	for{
//		time.Sleep(10000 * time.Second) 
//	}
//}
