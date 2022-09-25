package removeelement

func removeElement(nums []int, val int) int {
    
    i := 0 
    j := len(nums) - 1
    
    for i <= j  {
        if nums[i] == val {
            t := nums[j] 
            nums[i] = t 
            //don't move i 
            j--
        } else {
            i++ 
        }
        
    }
    
    return j + 1
    
}

