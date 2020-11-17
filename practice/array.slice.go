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


/*
Use make() first
- make: create value with slice, map, channel, it returns the value and use it at once
- new: request a piece of memory to be storge data, return the point(not init)

slice := make([]T, len, cap)
*/
func SliceCreate(){
    // Method1：建立一個包含数据的 string slice，适合用在知道 slice 里有什么元素时
    s1 := []string{"a", "b", "c", "d"}
    fmt.Printf("s1: %s\n", s1)
    // Method2：用 make 建立空的 slice，适合对 slice 中特定位置元素进行操作时
    s2 := make([]string, 4)
    fmt.Printf("s2: %s\n", s2)
    // Method3：空的 slice，一般搭配 append 使用
    var s3 []string
    fmt.Printf("s3: %s\n", s3)
    // Method4：大概知道需要多少元素时使用，搭配 append 使用
    s4 := make([]string, 0, 5)  
    fmt.Printf("s4: %s\n", s4)       // len=0 cap=5, []
}


// slice 可变更
func SliceCut() {
    s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
    fmt.Printf("s1 value: %d\n", s1)             // [1 2 3 4 5 6 7 8]
    fmt.Printf("s1 length: %d\n", len(s1))       // 8
    fmt.Printf("s1 capacity: %d\n", cap(s1))     // 8

    s2 := s1[3:6]
    fmt.Printf("s2 value: %d\n", s2)             // [4 5 6]
    fmt.Printf("s2 length: %d\n", len(s2))       // 3
    fmt.Printf("s2 capacity: %d\n", cap(s2))     // 5
}


// slice 扩容
func SliceAppend(){
    s1 := make([]string, 0, 10)
    fmt.Println(len(s1), cap(s1), s1)            // [] 0 10
    // s1[7] = "Hong"  -- fail   

    // slice expand(< 本体)
    s1 = append(s1, "Ming")
    fmt.Println(len(s1), cap(s1), s1)            // 1 10 [Ming]

    // slice expand(< 本体)
    s1 = s1[:8]
    s1[7] = "Hong"
    fmt.Println(len(s1), cap(s1), s1)            // 8 10 [Ming       Hong]

    // slice expand(> 本体)
    s2 := append(s1, make([]string, 20)...)
    fmt.Println(len(s2), cap(s2), s2)            // 28 28 [Ming       Hong                    ]
}


// zero value
func SliceZero(){
    var s []int
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)   // len=0 cap=0 []

    if s == nil {
        fmt.Println("nil!")  // nil!
    }
}


// range
// 如果只需要 index： for i := range pow
// 如果只需要 value： for _, value := range pwd
func SliceRange() {
    pow := []int {1, 2, 4, 8, 16, 32, 64, 128}

    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i , v)
    }
}