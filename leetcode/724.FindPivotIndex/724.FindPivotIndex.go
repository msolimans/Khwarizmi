package findpivotindex

func pivotIndex(nums []int) int {
	if len(nums) <= 1 {
		return 0 //index 0
	}

    //calc left and right 
	right := 0 
	left := 0 

	for i := 1; i <len(nums);i++ {
		right += nums[i]
	}

	for i := 0;i<len(nums);i++ {
		if left == right {
			return i
		}

		left += nums[i]
		
		if i < len(nums) - 1 { //edges 
			right -= nums[i+1] 
		}
	}

	return -1
	

}