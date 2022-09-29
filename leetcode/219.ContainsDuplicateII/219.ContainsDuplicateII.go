package containsduplicateii

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {

	//save into hashmap the indexes of each elem and calc abs difference to make sure it's less than or equal to k. if not, update index of elem 
	mp := map[int]int{}

	for i:=0;i<len(nums);i++ {
		if _,exists := mp[nums[i]]; exists {  
			if math.Abs(float64(i) - float64(mp[nums[i]])) <= float64(k) {
				return true 
			}
		} 		

		mp[nums[i]] = i	//we need to keep updating position of num[i] especially for cases like. 1,3,9,4,2,1,2,1
	}

	return false 

    
}