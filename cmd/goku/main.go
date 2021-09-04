
package main

import (
	"goku/internal/algo"
	"fmt"
)


func main(){

    arr := []int{10, 4, 41, 12, 32, 5, 22, 1, 6, 11}
    fmt.Println(arr)

    result := algo.QuickSort(arr, 0, len(arr)-1)
    fmt.Printf("result: %d\n", result)

}
