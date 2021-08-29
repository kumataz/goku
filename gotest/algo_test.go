package gotest

import (
	"testing"
	"goku/internal/algo"
	"fmt"
)

// fibonacci 
func TestFibonacci(t *testing.T)  {
	algo.Fibonacci(10)
}

// fibonacci closure way
func TestFibonacciClousure(t *testing.T){
	f := algo.FibonacciClosure()
	for i:=0; i<10; i++{
		fmt.Println(f())
	}
}