package checkifarrayissortedandrotated

func check(nums []int) bool {

	//we can get min element and then validate if everything after that element is greater than that element and everything before that element is also greater and come in order 

	//or we check every element in the array
	//in case next is greater => 
	
	
	
	// 5 6 7 1 2 3 4 
	
	if len(nums) <= 2 {
		return true 
	}


	//we now check 
	i := 1
	pivot := false 

	//I can use 2 approaches 
	//1) every element, the next is greater than itself except the pivot which is smaller than previous and next (only once)
	//2) find pivot index (left side should be greater than right side)

	//3 4 5 5 5 
	for i < len(nums) {
		if nums[i] < nums[i-1] && (i == len(nums) - 1 || nums[i] <= nums[i+1]) {
			if pivot { //only once is permitted, if this case exceeded once then we return false 
				return false 
				//5,6,3,4,3,4
			} else {
				pivot = true 
			}
		} else if nums[i] < nums[i-1] {
			return false 
			//5,4,3,2
		} 

		i++

		//otherwise, next is greater than or equal to current node 
	}


	if pivot { //in case there's pivot
		//last elem should be less than first elem in the array 
		if nums[len(nums) - 1 ] > nums[0] {
			return false 
		}
	}

	return true    
}