package basics

import(
    "errors"
    "fmt"
)


/* errors sample
Judge a number is positive or negative?  number == 0 is bug!
*/

func Positive(n int)(bool, error){
    if n == 0{
        return false, errors.New("Undefine!!!!")
    }
    return n > -1, nil
}

func Check(n int){
    pos, err := Positive(n)
    if err != nil{
        fmt.Println(n, err)
        return
    }
    if pos{
        fmt.Println(n, "is positive")
    }else{
        fme.Println(n, "is negative")
    }
}



/* Error Types
特点表现为：
- 一个指自定义 MyError{} 是一个 type，调用者可以用断言来转换这个类型获取更多的上下文信息。
- 能包装底层错误提供更多上下文
缺点为：
调用者要使用类型断言和类型 switch，就要让自定义的 error 变为 public。这种模型会导致和调用者产生强耦合，从而导致 API 变得脆弱。
结论：也尽可能避免使用 Error Type， 至少说避免将它们作为公共 API 的一部分
*/
type MyError struct{
    Msg string
    File string
    Line int
}

func (e *MyError) Error() string{
    return fmt.Sprintf("%s:%d:%s", e.File, e.Line, e.Msg)
}

func test()Error{
    return &MyError{"Something happened", "server.go", 42}
}

// Transfer
func ErrorTypes(){
    err := test()
    switch err := err.(type) {
    case nil:
        // success, nothing to do
    case *MyError:
        fmt.Println("error occurred on line:", err.Line)
    default:
        // unknow error
    }
}
