package majorityelement

func majorityElement(nums []int) int {
    
	//More's voting algorithm 

	//vote up or down based on occurence of elems 
	//we record current elem, then we keep it alive (output/result) as long as its vote is positive 
	//the winner is the one with the highest vote 
	res := nums[0]
	
	vote := 1 
	for i := 1;i<len(nums);i++ {
		//if current item is same as assumed result "vote up"
		if nums[i] == res {
			vote++ 
		} else { //vote down 
			vote-- 
		}

		//in case
		
		// once vote becomes zero, we may have a potential winner 
		if vote == 0 {
			res = nums[i] 
			vote = 1
		}

	}

	return res 
    
}