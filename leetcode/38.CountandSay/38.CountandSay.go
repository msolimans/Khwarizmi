package countandsay

import (
	"strconv"
	"strings"
)

//generating strings based on input
//12221111 => 113241
// 113241 => 2113121411
func countStr(s string, j int) string { 
    if j == 1 {
        return "1"
    }
    
    var sb strings.Builder 
    ch := s[0]
    cn := 1
    for i := 1; i< len(s);i++{
        if s[i] == ch {
            cn++  
        } else {
            sb.WriteString(strconv.Itoa(cn))
            sb.WriteByte(ch)
            ch = s[i]
            cn = 1
        }
    }
    sb.WriteString(strconv.Itoa(cn))
    sb.WriteByte(ch)
    
    return sb.String()
}
//dynamically recording results (1-30 based on constraints)
var arr = make([]string,31)


func count(s string, i, n int) string {
   
    ////once i get above n, return result of n 
    if i == n + 1 {
        return s
    }
    
	//calc current ith 
    ith := ""
    if arr[i] != "" {
        ith = arr[i]
    } else {
        ith = countStr(s, i)
        arr[i] = ith
    }
    
	//recursively pass ith result str   
    return count(ith, i+1,n)
}


func countAndSay(n int) string {
    
    return count("1", 1, n)
    
}