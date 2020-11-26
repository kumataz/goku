package gokupkg

import (
    "fmt"
    "net"
    "time"
    "hash/fnv"
)



/* 加密 encryption */

// net
func GetMacAddrs()(string,error){
      netInterfaces, err := net.Interfaces()
      if err != nil {
          return "",err
      }

     for _, netInterface := range netInterfaces {
          macAddr := netInterface.HardwareAddr.String()
              if len(macAddr) == 0 {
              continue
          }
         return macAddr,nil
      }
     return "",err
}


// hash output uint32
func Hash(s string) uint32{
        h := fnv.New32a()
        h.Write([]byte(s))
        return h.Sum32()
}



// add mac and time.now() as seed
func Hashseed() int64{
        mac_adr, _ := GetMacAddrs()
        t := time.Now().UnixNano() // int64
        return int64(Hash(fmt.Sprintf("%d %s",t,mac_adr)))
}


// Padding characters by a certain length
// example StrFormat32("kumata","0",8) --return--> 00kumata
func StrFormat32(raw ,sub string,slen int) string {
     switch {
     case len(raw) == slen:
         return raw
     case len(raw) > slen:
         return fmt.Sprintf("%.32s",raw[:slen])
     case len(raw) < slen:
         return fmt.Sprintf("%s%s",strings.Repeat(sub,slen-len(raw)),raw)
     default:
         /*should never meet here*/
         return raw
     }
 }
