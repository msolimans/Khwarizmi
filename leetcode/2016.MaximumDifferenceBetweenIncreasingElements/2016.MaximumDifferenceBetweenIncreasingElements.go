package maximumdifferencebetweenincreasingelements


func maximumDifference(nums []int) int {
    
    least := nums[0]
    max := -1 
    
    for i:=0;i<len(nums);i++ {
        if nums[i]<=least {
            least = nums[i]
        } else {//current elem is greater than least
            if  nums[i] - least  > max {
                max = nums[i] - least 
            }
        }
    }
    
    return max 
}

//5,4,3,4,3,2,3,4,5,6