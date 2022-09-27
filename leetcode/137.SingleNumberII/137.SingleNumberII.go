package singlenumberiii

import "sort"

 
func singleNumber(nums []int) int {
	sort.Ints(nums)
	
	for i :=0;i<len(nums);i++ {
		if i == len(nums)-1 || nums[i] != nums[i+1] {
			return nums[i]
		} else {
			i += 2 //jump three times in this case cause I know for sure at this point there are three same elems 
		}
	}

	return -1 
}