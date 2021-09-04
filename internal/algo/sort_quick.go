package algo

import(

)


/* 
快速排序，时间复杂度 O(nlogn)
分两个函数，一个找中心轴函数，一个是递归快排函数
*/

// 1. 快排
// 递归
func QuickSort(arr []int, left, right int) []int {
	if left < right{
		pivot := partition(arr, left, right)
		QuickSort(arr, left, pivot - 1)
		QuickSort(arr, pivot + 1, right)
	}

	return arr
}

// 2. 选取一个数作为中心轴，划分左右分区alice `low` and `high`
// 对数据元素进行左右调整，小于中心值移到左边，大于则移到右边
func partition(arr []int, low, high int) int {
	pivot := arr[low]
	for low < high{
		
		for low < high && arr[high] > pivot{
			high --
		}
		arr[low] = arr[high]

		for low < high && arr[low] < pivot{
			low ++
		}
		arr[high] = arr[low]
	}
	// after loop, low, high 重合
	arr[low] = pivot
	return low
}




