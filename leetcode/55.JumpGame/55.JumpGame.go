package jumpgame

func canJump(nums []int) bool {

	if len(nums) == 1 {
		return true 
	}

	//good index is the index from which I can jump to last index of the array 
	rightMostgoodIndex := len(nums) - 1
	i := len(nums) - 1

	for i >= 0 {
		if nums[i] + i >= rightMostgoodIndex {
			//update goodIndex 
			rightMostgoodIndex = i
		} 
		i--
	}
    
	//check if first index is a good index 
	return rightMostgoodIndex == 0
}