package practice

import(
    "fmt"
)

/*
# Record
1. array's length is immutable, and slice's length is variable.
2. "slice base on array", so array maybe include slice.
3. slice扩容时，golang会申请2倍的内存地址copy原有的过去，但如果长度大于1024时，则会调整为1.25倍，细节见runtime包中的slice.go：growslice():https://golang.org/src/runtime/slice.go?h=slice
4. slice缩容的话注意什么问题？切片缩容之后还是会引用底层的原数组，这有时候会造成大量缩容之后的多余内容没有被垃圾回收，可使用新建一个数组然后copy的方式。

- array: [3]string{"a", "b", "c"}
- slice: []string{"a", "b", "c"}

*/



func SliceDEMO() {
    a1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
    fmt.Printf("The length of a1: %d\n", len(a1))
    fmt.Printf("The capacity of a1: %d\n", cap(a1))
    fmt.Printf("The value of a1: %d\n", a1)
    // The length of a1: 8
    // The capacity of a1: 8
    // The value of a1: [1 2 3 4 5 6 7 8]

    s1 := a1[3:6]
    fmt.Printf("The length of s1: %d\n", len(s1))
    fmt.Printf("The capacity of s1: %d\n", cap(s1))
    fmt.Printf("The value of s1: %d\n", s1)
    // The length of s1: 3
    // The capacity of s1: 5
    // The value of s1: [4 5 6]
}

/*
Use make() first
- make: create value with slice, map, channel, it returns the value and use it at once
- new: request a piece of memory to be storge data, return the point(not init)
*/
func SliceAppend(){
    s2 := make([]int, 1024)
    fmt.Printf("s2  : len: %d, cap: %d\n", len(s2), cap(s2))

    s2e1 := append(s2, make([]int, 200)...)
    fmt.Printf("s2e1: len: %d, cap: %d\n", len(s2e1), cap(s2e1))

    s2e2 := append(s2, make([]int, 400)...)
    fmt.Printf("s2e2: len: %d, cap: %d\n", len(s2e2), cap(s2e2))

    s2e3 := append(s2, make([]int, 600)...)
    fmt.Printf("s2e3: len: %d, cap: %d\n", len(s2e3), cap(s2e3))

    // s2  : len: 1024, cap: 1024
    // s2e1: len: 1224, cap: 1280
    // s2e2: len: 1424, cap: 1696
    // s2e3: len: 1624, cap: 2048
}