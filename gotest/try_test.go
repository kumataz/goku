/*
go test 例子说明：
1. 文件名必须是 _test.go 结尾的，这样在执行`go test`的时候才会执行到相应的代码。
2. 必须import "testing" 这个包。
3. 所有的测试用例函数必须是Test开头
4. 测试用例会按照源代码中写的顺序依次执行
5. 测试函数TestXxx()的参数是`testing.T`，我们可以使用该类型来记录错误或者是测试状态
6. 测试格式：func TestXxx (t *testing.T),Xxx部分可以为任意的字母数字的组合，但是首字母不能是小写字母[a-z]，例如Testintdiv是错误的函数名。
7. 函数中通过调用testing.T的Error, Errorf, FailNow, Fatal, FatalIf方法，说明测试不通过，调用Log方法用来记录测试的信息。
*/

package gotest

import (  
  // "fmt"
  "testing"
)

func TestSum(t *testing.T) {  
  numbers := []int{1, 2, 3, 4, 5}
  expected := 15
  actual := Sum(numbers)

  // compare test
  if actual != expected {
    t.Errorf("Expected the sum of %v to be %d but instead got %d!", numbers, expected, actual)
  }
}
