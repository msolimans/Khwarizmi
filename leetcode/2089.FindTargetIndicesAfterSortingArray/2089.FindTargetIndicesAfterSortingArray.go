package findtargetindicesaftersortingarray

import "sort"

func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)

	targetIndex := findTarget(nums, target)
	if targetIndex == -1 {
		//not found 
		return []int{}
	}

	i, j := expand(nums, target, targetIndex)
	res := []int{}
	for x := i;x <= j;x++ {
		res = append(res, x)
	}

	return res 
}


func expand(nums []int, target, indx int) (int,int) {
	i, j := indx, indx 

	for i >= 0 && nums[i] == target {
		i--
	}
	//we move forward as long as current element equals to target (last one will not be same thu)
	for j <= len(nums) - 1 && nums[j] == target {
		j++ 
	}

	//we subtract i and j (-1) as the last element is not same as target 
	return i+1,j-1

}

func findTarget(nums []int, target int) int {
	s, e := 0,len(nums) - 1 

	for s <= e {
		mid := ( s + e ) / 2 

		if target == nums[mid]  {
			return mid 
		} else if nums[mid] < target {
			s  = mid + 1 
		} else {
			e = mid - 1
		}
	}

	return -1 
}