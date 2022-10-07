package zigzagconversion

import "strings"

func convert(s string, numRows int) string {
    
	if numRows == 1 {
		return s 
	}
	
    sbs := make([]strings.Builder, numRows) 
    
    i := 0 
    dir := 1
    
    for _, c := range s  {
        //rune  
        sbs[i].WriteRune(c)
    
        // i++ (if we are at `numRows - 1` flip dir and make it operation --) 
		//or i-- if we are at zero, flip operation and switch to ++ 
        if dir == 1{//down 
            if i == numRows - 1 {
                i--
                dir = -1 
            } else {
                i++    
            }
        } else {//up 
            if i == 0 {
                i++ 
                dir = 1
            } else {
                i--
            }   
        }
    }
    
    var res strings.Builder 
    for i := 0 ;i<len(sbs);i++ {
        res.WriteString(sbs[i].String())
    }
    return res.String()
    
}