package singlenumberiii

import "sort"

func singleNumber(nums []int) []int {
	mp := map[int]int{} 
	for _,n:= range nums {
		mp[n] = mp[n] + 1 
	}

	res := []int{}
	for k,v := range mp {
		if v != 2 {
			res = append(res, k)
		}
	}

	return res 
}


//2nd soln 
func singleNumber2(nums []int) []int {
	sort.Ints(nums)
	res := []int{}
	for i :=0;i<len(nums);i++ {
		if i == len(nums)-1 || nums[i] != nums[i+1] {
			res = append(res, nums[i])
		} else {
			i++ //jump twice 
		}
	}

	return res 
}