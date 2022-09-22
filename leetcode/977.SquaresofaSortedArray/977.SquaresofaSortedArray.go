package squaresofasortedarray

import "math"

func sortedSquares(nums []int) []int {
    res := make([]int, len(nums))

	//pointer in the result array 
	index := len(res)  - 1 
	//negatives in the left side of the array can be greater than positives in the most right side of the array 
	//2 pointers to know which one is greater 
	i := 0 
	j := len(nums) -  1

	for i <= j {
		//think greedy here (who is bigger then take it and put it in the back of the array)
		if math.Abs(float64(nums[i])) > math.Abs(float64(nums[j])) {
			res[index] = nums[i] * nums[i]
			i++
		} else {
			res[index] = nums[j] * nums[j]
			j--
		}

		index--
	}
	
	return res 
	
}