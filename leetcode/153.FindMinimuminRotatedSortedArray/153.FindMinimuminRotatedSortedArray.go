package findminimuminrotatedsortedarray

func findMin(nums []int) int {
    
    
    s := 0 
    e := len(nums) - 1 
    n := len(nums) 
    
    for s <= e {
        //we need to see where the pivot might be 
        mid := (s+e) / 2
        next := (mid+1)%n // to make sure we don't go outside array boundary 
        prev := (mid-1+n)%n
        
        if nums[mid] <= nums[prev] && nums[mid] <= nums[next] {
            //this is the target 
            return nums[mid]
        } else if nums[mid] <= nums[e] {//we are in the first half of array, go right 
            e = mid - 1
        } else { // in the right side of array, go left 
            s = mid + 1 
        }
    }
    
    return -1 
}


////////////////////////////////////////////////////////////////////////////////

//11 13 15 17 => (i-1+n)%n
//4 5 6 1 (pivot might be at the end of array, less than previous)

//pivot is the one that is smaller than next and previous 
//     5 
//   4 
// 3 
//         2
//       1