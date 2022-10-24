package poweroffour

//sqrt(n)
func isPowerOfFour(n int) bool {
    if n == 1 {
        return true 
    }
    
    if n < 4 {
        return false 
    }
    
    if n % 4 != 0 {
        return false 
    }
    
    return isPowerOfFour(n/4)
}