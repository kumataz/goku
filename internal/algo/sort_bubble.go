package algo

import(

)


// 冒泡排序 O(n^2)
// 输入一个待排序数组，输出排序后的数组
// arr := []int{10, 4, 41, 12, 32, 5, 22, 1, 6, 11}
func BubbleSort(arr []int) []int{
	for i:=0;i<len(arr)-1;i++{
		for j:=0;j<len(arr)-1-i;j++{  // kp
			if arr[j] > arr[j+1]{
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}


// 改进版本 O(n^2), 加入flag判断是否需要进入二层嵌套, 减少循环次数
func BubbleSortOpt1(arr []int) []int{
	for i:=0;i<len(arr)-1;i++{
		flag := false
		for j:=0;j<len(arr)-1-i;j++{ 
			if arr[j] > arr[j+1]{
				flag = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if !flag {
			break
		}
	}
	return arr
}