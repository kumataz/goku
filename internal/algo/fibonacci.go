package algo

import (
	"fmt"
)

// fibonacci without closure
func Fibonacci(num int) {
	b1 := 1
	b2 := 0
	bc := 0
	for i:=0; i<num; i++{
		bc = b1 + b2
		b1 = b2
		b2 = bc 
		fmt.Println(bc)
	}
}

/*
常规做法的缺点：
1. 把循环的次数交给了`Fibonacci`函数，每次函数结束后再执行都是从b1=1开始; (效率低)
2. 如果想继续下去必须再函数外声明变量 (变量泛滥)
闭包：
缩小变量作用域，减少对全局变量的污染
*/

// fibonacci closure <闭包>
func FibonacciClosure() func() int {
	b1 := 1
	b2 := 0
	bc := 0
	return func() int{
		bc = b1 + b2
		b1 = b2
		b2 = bc 
		return bc
	}
}



