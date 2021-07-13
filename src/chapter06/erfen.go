package main

// 二分查找实现代码
//func binarySearch(nums []int, num int, low int, high int) int {
//	// 递归终止条件
//	if low > high {
//		return -1
//	}
//	// 通过中间元素进行二分查找
//	mid := (low + high) / 2
//	// 递归查找
//	if num > nums[mid] {
//		// 如果待查找数据大于中间元素，则在右区间查找
//		return binarySearch(nums, num, mid+1, high)
//	} else if num < nums[mid] {
//		// 如果待查找数据小于中间元素，则在左区间查找
//		return binarySearch(nums, num, low, mid-1)
//	} else {
//		// 找到了，返回索引值
//		return mid
//	}
//}
//func main() {
//	nums := []int{4, 6, 5, 3, 1, 8, 2, 7}
//	sort.Ints(nums) // 先对待排序数据序列进行排序
//	fmt.Printf("Sorted nums: %v\n", nums)
//	num := 5
//	index := binarySearch(nums, num, 0, len(nums)-1)
//	if index != -1 {
//		fmt.Printf("Find num %d at index %d\n", num, index)
//	} else {
//		fmt.Printf("Num %d not exists in nums\n", num)
//	}
//}

// 二分查找变形版本1：查找第一个值等于给定值的元素
//func binarySearchV1(nums []int, num, low, high int) int {
//	if low > high {
//		return -1
//	}
//	mid := (low + high) / 2
//	if num < nums[mid] {
//		return binarySearchV1(nums, num, low, mid-1)
//	} else if num > nums[mid] {
//		return binarySearchV1(nums, num, mid+1, high)
//	} else {
//		// 除了需要判断第一个与 num 相等的元素索引外，其他和普通二分查找逻辑一致
//		if mid == 0 || nums[mid-1] != num {
//			return mid
//		} else {
//			return binarySearchV1(nums, num, low, mid-1)
//		}
//	}
//}
//func main() {
//	nums := []int{1, 2, 3, 3, 4, 5, 6}
//	num := 3
//	index := binarySearchV1(nums, num, 0, len(nums)-1)
//	if index != -1 {
//		fmt.Printf("Find num %d first at index %d\n", num, index)
//	} else {
//		fmt.Printf("Num %d not exists in nums\n", num)
//	}
//}

// 二分查找变形版本3：查找第一个大于等于给定值的元素
//func binarySearchV3(nums []int, num, low, high int) int {
//	if low > high {
//		return -1
//	}
//	mid := (low + high) / 2
//	if num <= nums[mid] {
//		// 判断条件变更，找到第一个大于等于给定值的元素
//		// 最左侧元素值，或者当前 mid 位置前一个元素值小于给定值
//		if mid == 0 || nums[mid-1] < num {
//			return mid
//		} else {
//			return binarySearchV3(nums, num, low, mid-1)
//		}
//	} else {
//		return binarySearchV3(nums, num, mid+1, high)
//	}
//}
//func main() {
//	nums := []int{1, 2, 3, 3, 4, 5, 6}
//	num := 3
//	index := binarySearchV3(nums, num, 0, len(nums)-1)
//	if index != -1 {
//		fmt.Printf("Find first num greater or equal than %d at index %d\n", num, index)
//	} else {
//		fmt.Printf("Num %d not exists in nums\n", num)
//	}
//}

// 二分查找变形版本4：查找最后一个小于等于给定值的元素
//func binarySearchV4(nums []int, num, low, high int) int {
//	if low > high {
//		return -1
//	}
//	mid := (low + high) / 2
//	if num >= nums[mid] {
//		// 判断条件变更，找到最后一个小于等于给定值的元素
//		// 最右侧元素值，或者当前 mid 位置后一个元素值大于给定值
//		if mid == len(nums)-1 || nums[mid+1] > num {
//			return mid
//		} else {
//			return binarySearchV4(nums, num, mid+1, high)
//		}
//	} else {
//		return binarySearchV4(nums, num, low, mid-1)
//	}
//}
//func main() {
//	nums := []int{1, 2, 3, 3, 4, 5, 6}
//	num := 3
//	index := binarySearchV4(nums, num, 0, len(nums)-1)
//	if index != -1 {
//		fmt.Printf("Find last num less or equal than %d at index %d\n", num, index)
//	} else {
//		fmt.Printf("Num %d not exists in nums\n", num)
//	}
//}
