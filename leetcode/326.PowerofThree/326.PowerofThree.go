package powerofthree

import "math"


func isPowerOfThree(n int) bool {
    if n == 0 {
        return false 
    }
    if n == 1 {
        return true 
    }
    
    if n % 3 > 0 {
        return false 
    }
    
    return isPowerOfThree(n/3)
}

//the max int is 3^19
// 3^19 % n == 0 => it is power of 3 otherwise it is not
func isPowerOfThree2(n int) bool {
    if n <= 0 {
        return false 
    }

	//max we can get 
	max := math.Pow(3,19)
	//if this max number is divisible by n we return true otherwise false 
	return int(max) % n  == 0 
}