package happynumber

func isHappy(n int) bool {

	for n/10 != 0 { //continue as long as we are higher than 10 (base case is between 1 - 9 there are only 2 numbers that are happy ( 1 or 7) )
		n = calcSum(n)
	}
	
	//we can use hashmap instead if we don't even know the base case scenario here 
	return n == 1 || n == 7 

}

func calcSum(n int) int { 
	//calc sum of a value n
	sum := 0 
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}

	return sum 
}

////////////////////////////////////////
//2nd solution based on recursion and hashmap (uses extra space)

var visited = map[int]bool{}
func isHappy2(n int) bool {
	
	if n == 1 {
		return true 
	}

	if _, exists := visited[n]; exists { 
		return false 
	}
	
	visited[n] = true 
	n = calcSum(n)
	return isHappy2(n)
}
