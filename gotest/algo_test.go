package gotest

import (
	"testing"
	"goku/internal/algo"
	"fmt"
)



func TestAlog(t *testing.T){
	arr := []int{12, 2, 52, 41, 4, 7, 15, 56, 10, 51}
	result := algo.SearchOrder(arr, 7)
	fmt.Printf("result: %d\n", result)
}