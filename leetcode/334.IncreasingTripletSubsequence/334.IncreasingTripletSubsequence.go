package increasingtripletsubsequence

import "math"

func increasingTriplet(nums []int) bool {
    //lets assume first second and third are the numbers 
    //or i j k 
    if len(nums) < 3 {
        return false 
    }
    
    i,j := math.MaxInt,math.MaxInt
    for k := 0;k < len(nums);k++ {
        //if current number is greater than previous 2 elements then return true 
        if nums[k] > i && nums[k] > j  {
            return true
        } else if nums[k]>i && nums[k] <j {//in between previous 2 elems [i and j]            
            //if nums[k] is greater than i and less than j then it is clear that we find j in i < j < k
            j = nums[k] //keep i and update j 
        } else if nums[k] < i {
            i = nums[k] //less than i, only update i 
        }
    }
    
    return false 
}