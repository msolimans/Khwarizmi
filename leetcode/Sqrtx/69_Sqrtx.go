package sqrtx

func mySqrt(x int) int {
	//use binary search for that since all integers are all sorted to find the correct number that multiplies in itself and gives x 
	i, j := 0, x 

	for i < j {
		mid := (i+j+1) / 2
		res := mid * mid 
		if res == x {
			return mid 
		} else if res < x {
			//move towards bigger digits 
			i = mid 
		} else { //bigger (lower range is needed)
			j = mid - 1
		}
	}

	return i //return the best one that can be closer to x (cases like 8 while 3*3 = 9 but 2*2 = 4 which is the best match)
}
