package rotatearray

import "fmt"

func rotate(nums []int, k int)  {
	//[1,2,3,4,5,6,7]
	//[      1,2,3,4]      

	k %= len(nums)

	ri := (len(nums) - k) 
	rk := make([]int, k)
	i := 0
	for ;ri<len(nums);ri++ {
		rk[i] = nums[ri]
		i++ 
	}
	fmt.Println(rk) 

	in := len(nums) - 1
	jn := len(nums) - k - 1

	for jn >= 0 {
		nums[in] = nums[jn]
		jn--
		in--
	}

	fmt.Println(nums)

	for i := 0;i<k;i++ {
		nums[i] = rk[i]
	}
}