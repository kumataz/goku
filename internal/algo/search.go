package algo

import(
	"fmt"
)

/*
查找算法，最常见的就是两种：顺序查找和二分法查找
*/

// 1.顺序查找， 时间复杂度 = O(n)
// Q: 提供一个数组array，在这个数组里找出key=7的值返回其index，没有则返回-1
// arr := []int{12, 2, 52, 41, 4, 7, 15, 56, 10, 51}
func SearchOrder(arr []int, key int) int{
	index := -1

	for i, value := range arr {
		if value == key {
			index = i
			break
		}
	}
	if index != -1{
		fmt.Printf("success find value <%d> 's index:", key)
		return index
	}else{
		fmt.Printf("cannot find value <%d> 's index", key)
		return index
	}
}


// 2. 二分查找