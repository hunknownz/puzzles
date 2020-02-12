package removeDuplicatesFromSortedArray

func removeDuplicates(nums []int) int {
	newNumI, numsLen := 0, len(nums)
	if numsLen == 0 {
		return 0
	}
	for i:=1; i<numsLen; i++ {
		if nums[i] != nums[i-1] {
			newNumI = newNumI + 1
			nums[newNumI] = nums[i]
		}
	}
	return newNumI + 1
}
