package algo

import(
	"fmt"
)

/*
查找算法，最常见的就是两种：顺序查找和二分法查找
*/

// 1.顺序查找， 时间复杂度 = O(n),  查找表类型：无序表
// Q: 提供一个数组array，在这个数组里找出key=10的值返回其index，没有则返回-1
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


// 2. 二分查找， O(logn), 查找表类型：有序表
// Q: 提供一个数组array，在这个数组里找出key=10的值返回其index，没有则返回-1
// arr := []int{0, 3, 3, 5, 6, 7, 7, 8, 9, 10, 14, 15}
func SearchBinry(arr []int, key int) int{

	if len(arr) == 0{
		return -1
	}

	left := 1
	right := len(arr)-1
	mid := 0

	for i:=0; i<len(arr); i++{
		if key == arr[mid] || left >= right{
			break
		}else if key < arr[mid]{
			right = mid - 1
			mid = (left + right) / 2
		}else if key > arr[mid]{
			left = mid + 1
			mid = (left + right) / 2
		}
	}

	if key != arr[mid]{
		mid = -1
	}

	return mid
}