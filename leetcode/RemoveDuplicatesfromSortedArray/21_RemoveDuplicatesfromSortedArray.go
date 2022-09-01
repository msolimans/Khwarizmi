package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	//i is a quicker pointer while head is always pointing to the last element where there is no duplicate 
	head, i := 0, 1
	for i < len(nums) {
		if nums[i] == nums[head] {
			i++
		} else {
			//not equal 
			head++ //move to next spot for new elem to come in 
			nums[head] = nums[i]
		}
	}

	return head + 1
}
