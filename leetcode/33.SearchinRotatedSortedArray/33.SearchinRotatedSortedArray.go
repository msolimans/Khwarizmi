package searchinrotatedsortedarray

func search(nums []int, target int) int {
    
    //we can use brute force can be o(n) .. search n all items 
    
    s := 0
    e := len(nums) - 1
    
    for s <= e { 
        mid := (s+e) / 2 
        
        if nums[mid] == target {
            return mid 
        } else if nums[mid] >= nums[s] {//we are in the left side 
            //now check where target might be between start and mid 
            if target >= nums[s] && target <= nums[mid] { //between mid and start 
               e = mid - 1
            }  else {//between mid and pivot point 
                s = mid + 1 
            }
        } else { // we are in the second half of array 
            if target >= nums[mid] && target <= nums[e] { //between mid and end of array 
                s = mid + 1
            } else { //between mid and pivot 
                e = mid - 1
            }
            
        }
    
    }
    
    return -1 
    
}


//12345678
//BS into sorted array to get right index 

//       4
//     3
//   2
// 1 

//            4
// a[0]    3    
//                2   a[n-1]
//             1 