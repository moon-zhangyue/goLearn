package main

//
//import "fmt"
//
////插入排序
//func insertionSort(nums []int) []int {
//	if len(nums) <= 1 {
//		return nums
//	}
//
//	for i := 0; i < len(nums); i++ {
//		//每次从未排序区间取一个数据 value
//		value := nums[i]
//		//在已排序区间找到插入位置
//		j := i - 1
//		for ; j >= 0; j-- {
//			//如果比value大后移
//			if nums[j] > value {
//				nums[j+1] = nums[j]
//			} else {
//				break
//			}
//		}
//		//插入数据value
//		nums[j+1] = value
//	}
//
//	return nums
//}
//
//func main() {
//	nums := []int{2, 3, 5, 1, 7, 8, 6}
//	nums = insertionSort(nums)
//	fmt.Println(nums)
//}
