package findclosestnumbertozero

import "math"

func findClosestNumber(nums []int) int {
    res := nums[0]
    dis := math.Abs(float64(nums[0] - 0))

    for i := 1;i<len(nums);i++ {
        //calc distance between current elem and zero  
        c := math.Abs(float64(nums[i] - 0))
        //if its equal, compare previous stored result vs current result and update result with the bigger value 
        if c == dis {
            if nums[i] > res {
                res = nums[i]
            }
        } else if c < dis {
            //current distance is less, update result and its associated distance 
            dis = c 
            res = nums[i]
        }
    }

    return res 
}