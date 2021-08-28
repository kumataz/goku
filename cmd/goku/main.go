
package main

import (
	"goku/internal/algo"
	"fmt"
)



func main(){
	arr := []int{12, 2, 52, 41, 4, 7, 15, 56, 10, 51}
	result := algo.SearchOrder(arr, 7)
	fmt.Printf("result: %d\n", result)
}
