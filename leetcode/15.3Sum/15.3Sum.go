package sum

import (
	"fmt"
	"sort"
)
func threeSum(nums []int) [][]int {
    
    sort.Ints(nums)
    res := [][]int{}
    for i := 0;i<len(nums);i++ {
        goal := 0 - nums[i]
        twos := twoSum(nums,i+1,goal)
        for _,v := range twos {
            res = append(res, []int{nums[i], v[0],v[1]})
        }
    }
    
    //dups
    mp := map[string]bool {}
    fres := [][]int{}
    for _,v := range res {
        if _,exists := mp[fmt.Sprintf("%v,%v,%v",v[0], v[1],v[2])]; !exists {
            mp[fmt.Sprintf("%v,%v,%v",v[0], v[1],v[2])] = true 
            fres = append(fres, v)
        }
    }
    
    return fres
}

func twoSum(nums []int, s, k int) [][]int {
    i, j := s, len(nums) - 1
    
    res := [][]int{}
    
    for i < j {
        if nums[i] + nums[j] == k {
            res = append( res, []int{nums[i], nums[j]})
            //continue searching as we may have other res //-2,0,1,1,2
            i++
            j--
        } else if nums[i] + nums[j] < k {
            i++
        } else {
            j--
        }
    }
    
    return res
}

