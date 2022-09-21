package validperfectsquare

func isPerfectSquare(num int) bool {
	if num == 0 || num == 1 {
		return true
	}

	i, j := 1,num
	for i<j {
		mid := (i+j+1)/ 2
		res := mid * mid 
		if res == num {
			return true
		} else if res < num {
			i = mid 
		} else {
			j = mid - 1 //exclude mid since its higher than num 
		}


	}

	return false 
    
}