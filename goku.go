
package main

import(
    "fmt"
    "errors"
    // . "goku/basics"
)



func EchoEmpty(request string)(response string, err error){
    if request == ""{
        err = errors.New("Can not empty!")
        return
    }
    response = fmt.Sprintf("Echo: %s\n", request)
    return
}



func main() {
    for _, req := range []string{"", "hello!"} {
      fmt.Printf("request: %s\n", req)
      resp, err := Echo(req)
      if err != nil {
        fmt.Printf("error: %s\n", err)
        continue
      }
      fmt.Printf("response: %s\n", resp)
    }
}
