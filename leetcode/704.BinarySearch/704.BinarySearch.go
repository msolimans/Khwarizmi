package binarysearch

func search(nums []int, target int) int {
    s := 0
    e := len(nums) - 1
    
    for s <= e {
        mid := (s + e)/2
        if nums[mid] == target {
            return mid 
        } else if nums[mid] < target {
            s = mid + 1
        } else {//greater than 
            e = mid - 1
        }
    }
    
    return -1 
}