package climbingstairs

var m = map[int]int{} 

func climbStairs(n int) int {
    //calc where I am comes from previous steps 
	//since we only have 2 possible moves 1 or 2: we can depend on previous 2 steps -
	// if I am at step 3, we can only come from step 1 or 2  

	//step 1 -> only one way (jump 1 step) 
	//step 2 -> only 2 ways (jump 1 + 1 step OR jump 2 steps)
	//step 3 -> now you can combine whatwever we got from step 1 and step 2 which will result in 3 ways 
	if n == 1 {
		return 1 
	} else	if n == 2 {
		return 2
	} else {
		if _, exists := m[n - 1]; !exists {
			m[n-1] = climbStairs(n - 1)
		} 
		if _, exists := m[n - 2]; !exists {
			m[n-2] = climbStairs(n - 2)
		} 

		return m[n - 1] + m[n - 2]
	}
}