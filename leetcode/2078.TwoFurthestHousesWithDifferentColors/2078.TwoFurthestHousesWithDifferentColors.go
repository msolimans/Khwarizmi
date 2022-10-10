package twofurthesthouseswithdifferentcolors

import "math"

func maxDistance(colors []int) int {
   
    //get the last house that's different than the first house 
    j := len(colors) - 1 
    for j >0 && colors[j] == colors[0] {
        j--
    }
    
    //get the first house that's different than the last house 
    // len(colors)-1-i (to get the maxDiff from right)
    i := 0 
    for i < len(colors) && colors[i] == colors[len(colors) - 1] {
        i++ 
    }
    
    
    //get the max difference 
    return int(math.Max(float64(j), float64(len(colors)-1-i)))
}

// 1,8,3,1,8,3,9
//             ^j 
// i

// 8,8,8,8,8,1,2,3
// 
// 8,8,1,2,3,4,5,8,8,8,8
//     i
//             j
              