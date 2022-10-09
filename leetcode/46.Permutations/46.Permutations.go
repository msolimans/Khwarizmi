package permutations


var res [][]int = [][]int{}

func helper(nums []int, j int) {
    // //1             2                 3 
    // 1 23           2 13                3 21
    // 12 [3]         21 [3]              32 [1]
	//    j=2  			 j=2				 j=2
    // 13 [2]         23 [1]              31 [2]
	// 	   j=2			  j=2				 j=2
	if j == len(nums) - 1 { //stopping condition is last elems in the array 
		// fmt.Println(j,nums) /j is always the last elem in the array 
		l := append([]int(nil), nums...)
		// l := make([]int, len(nums))
		// copy(l,nums)
		res = append(res, l)// we append whatever in the array 
	}
	
	for i := j; i<len(nums) && j<len(nums);i++ {  
		swap(nums, i, j) 
        helper(nums, j+1) //next to whatever we passed 
        swap(nums,i,j) 
    }
}

func swap(nums []int, i,j int ) {
    t := nums[i]
    nums[i] = nums[j]
    nums[j] = t 
}

func permute(nums []int) [][]int {
    res = [][]int{} //init res 
    helper(nums,0) //we start with zero-indexed node
    return res 
}