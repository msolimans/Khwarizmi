package findfirstandlastpositionofelementinsortedarray

func searchRange(nums []int, target int) []int {
	s := 0
	e := len(nums) - 1

	for s <= e {
		mid := (s + e) / 2 

		if nums[mid] == target {
			return  expand(nums, target, mid)
		} else if nums[mid] < target {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}

	//in case not found 
	return []int{-1,-1}
}

//return smallest and largest index that has target elem
func expand(nums []int, target, index int) []int { 

	i, j := index, index

	for i >= 0 && target == nums[i] {
		i--
	}

	for j < len(nums) && nums[j] == target {
		j++
	}

	return []int{i +1, j-1}
}