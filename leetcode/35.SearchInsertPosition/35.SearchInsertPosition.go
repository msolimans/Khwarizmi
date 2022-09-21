package searchinsertposition

func searchInsert(nums []int, target int) int {
    //special case handling for leetcode (it assumes in case nums: [1] and target: 1, it expects 0)
    if len(nums) == 1 && nums[0] == target {
        return 0
    }
    
    if target > nums[len(nums) - 1] {
        return len(nums)
    } else if target <= nums[0] {
        return 0 
    } 

    s := 0
    e := len(nums) - 1 
    mid := 0
    
    for s <= e {
        mid = (s + e) / 2
        
        if nums[mid] == target {
            //find the item using binary search. in case found return it 
            return mid 
        } else if nums[mid] < target {
            s = mid + 1
        } else {
            e = mid - 1
        }
        
    }
    

    //if not found above, check if target is bigger than the closest found item then return the next index 
    if target > nums[mid]  {
        return mid + 1
    } else {
        //otherwise (less than or equal to target) return same index as array elements should be shifted to leave a room for that target to be inserted
        return mid 
    } 
    
    
}