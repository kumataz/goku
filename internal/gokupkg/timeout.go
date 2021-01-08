package gokupkg

import (
	"fmt"
	"time"
)

func main() {
	// loop 5 s
	timeout := time.After(time.Second * 5)
	finish := make(chan bool)

	go func() {
		for {
			select {
			case <-timeout:
				// time out message
				fmt.Println("timeout")
				finish <- true
				return
			default:

				// loop in it
				// loooooooooooop ~
				for{
					fmt.Printf("go~~ %d\n", count)
				}

			}
			// ouput 1 round/s
			time.Sleep(time.Second * 1)
		}
	}()

	<-finish

	fmt.Println("done!")
}
