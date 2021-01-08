
package basics

import(
    "fmt"
)

/*
Map: Key-Value data struct in golang
*/

func MapCreate() {
    aMap := map[int]string{
        1: "one",
        2: "two",
        3: "three",
    }

    k := 2
    v, ok := aMap[k]
    if ok{
        fmt.Printf("%d: %s\n", k, v)
    }else{
        fmt.Println("k not found!")
    }
}

/*
// key'type cannot be "map, slice, channel, func"

var badMap2 = map[interface{}]int{
  "1":   1,        // pass
  []int{2}: 2,     // 这里会引发panic。
  3:    3,         // pass
}
*/
