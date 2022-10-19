package countingbits

//naive solutions
func countBits(n int) []int {
    res := make([]int, n+1)
    for i := 0;i<len(res);i++ {
		//call count and record in res[i]
        res[i] = count(i)
    }
    
    return res 
}

//count how many 1s in a number 
func count(n int) int {
    c := 0 
    for i := 0;i < 32;i++ {
        m := 1 << i
        if m & n != 0 {
            c++
        }
    }
    return c 
}

//better solution (dynamic programming)
//based on previous results, calc current value 
//at 2, 4, 8, 16, ..etc => it is 1 
//in between value of i-pow  + 1 
func countBits2(n int) []int {
    //dynamic programming 
    //range 2-3 => 4-7 => 8 - 15 => 16 - 31 
    res := make([]int, n+1 )
    
    pow2 := 1
    
    for i := 1;i<len(res); i++ { 
        if i == pow2*2 { 
            res[i] = 1 //it is only one bit here 
            pow2 = pow2 * 2 //what is next 
        } else { //if we haven't passed (pow2*2) 
            res[i] = 1 + res[i-pow2] // 13 - 8 or 14 - 8 
        }
    }
    
    return res 
} 