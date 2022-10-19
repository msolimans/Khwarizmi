package poweroftwo

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false 
	}
	return n & (n - 1) == 0
}

func isPowerOfTwo2(n int) bool {
	if n == 1 {
        return true 
    }
	
    if n % 2 > 0 {
        return false     
    }
	
    return isPowerOfTwo(n/2)
}