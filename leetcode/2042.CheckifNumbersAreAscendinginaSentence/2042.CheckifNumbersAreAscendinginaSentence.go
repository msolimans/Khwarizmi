package checkifnumbersareascendinginasentence

import "math"

func areNumbersAscending(s string) bool {
    prev := math.MinInt32
    i := 0 
    for i < len(s) {
        if s[i] == ' '  {
            i++
            continue 
        }

        if !isNumber(s[i]) {
            i++ 
            continue
        }
        
        curr, newIndex := extractNumber(s,i)
        if curr <= prev {
            return false 
        }
        prev = curr
        i = newIndex //jump 
    }

    return true 
    
}
func isNumber(c byte) bool {
    return int(c - '0') >= 0 && int(c - '0') <= 9
}

func extractNumber(s string, i int) (int, int) { //returns str as int and new index
    res := 0 
    for i < len(s) && isNumber(s[i]) {
        res = res*10 + int(s[i] - '0')
        i++
    }
    return res, i
}
