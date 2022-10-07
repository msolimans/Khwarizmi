package rotatestring

func rotateString(s string, goal string) bool {
    
    //brute force 
    //keep shifting n times 
    c := len(s) 
    for c > 0 {
        s = shiftString(s)
        if s == goal {
            return true 
        }
        
        c--
    }
    
    return false 
}


func shiftString(s string) string {
    return string(s[len(s)-1])+s[0:len(s)-1]
    //12345
    //51234
    
}
