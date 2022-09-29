package containsduplicate

func containsDuplicate(nums []int) bool {
    mp := map[int]bool{}
    
    for _, n := range nums {
        if _,exists := mp[n]; exists {
            return true 
        }
        mp[n] = true 
    }
    
    return false
    
}