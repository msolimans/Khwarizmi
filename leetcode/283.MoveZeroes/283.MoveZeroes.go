package movezeroes

//maintains 2 pointer to keep track of beginning of zeros and beginning of non-zeros
func moveZeroes(nums []int)  {

	izero := findFirstZero(0, nums)
	inon := findFirstNonZero(0,nums) 

	for izero != -1 && inon != -1 {

		//if nonzero is before zero, keep moving nonzero index
		if inon < izero {
			inon = findFirstNonZero(inon+1,nums) 	//find none-zero after current index 
			continue
		}

		//swap 
		nums[izero] = nums[inon]	
		nums[inon] = 0
		
		//update indexes
		izero = findFirstZero(izero+1, nums)//find zero after current index 
		inon = findFirstNonZero(inon+1,nums) 	//find none-zero after current index 

	}
    
}

func findFirstZero(i int, nums []int) int {
	for ;i<len(nums);i++ {
		if nums[i] == 0 {
			return i 
		}
	}

	//greater than length (meaning it doesn't exist)
	return -1 
}


func findFirstNonZero(i int, nums []int) int {
	for ;i<len(nums);i++ {
		if nums[i] != 0 {
			return i 
		}
	}

	return -1
}


//////

//2nd solutions 
func moveZeroes2(nums []int)  {

	 
	for i :=0 ;i<len(nums) - 1;i++ {
		if nums[i] != 0 {
			continue 
		}

		//if found zero, keep moving until you see non-zero and then swap it with that current zero-index
		//only donwside here is number of iterations we repeat here in cases like 
		//1,00000000000000000000000000000000000000,0,0,2,3,4
		j := i 
		for ;j<len(nums) - 1;j++ {
			if nums[j] != 0 {
				break
			}
		}
		//do swap 
		nums[i] = nums[j] 
		nums[j] = 0
	}
    
}

