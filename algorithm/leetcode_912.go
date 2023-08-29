package main

/*
*
912. 排序数组 虽然跟java不一样 引用对象要传地址啊 我都改成地址了 为啥你还是跑失败了
*/
func sortArray(nums []int) []int {
	quicksort(&nums, 0, len(nums)-1)
	return nums
}

func partition(arr *[]int, l, r int) int {
	nums := *arr
	pivot := nums[l]
	index := l + 1
	for i := index; i <= r; i++ {
		if nums[i] < pivot {
			//swap(arr, index, i)
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}
	//swap(arr, l, index-1)
	nums[l], nums[index-1] = nums[index-1], nums[l]
	return index - 1
}

func quicksort(nums *[]int, l, r int) {
	if l < r {
		index := partition(nums, l, r)
		quicksort(nums, l, index-1)
		quicksort(nums, index+1, r)
	}
}

//func swap(arr *[]int, a, b int) {
//	nums := *arr
//	temp := nums[a]
//	nums[a] = nums[b]
//	nums[b] = temp
//}

func main() {
	nums := []int{5, 2, 3, 1}
	sortArray(nums)
}
