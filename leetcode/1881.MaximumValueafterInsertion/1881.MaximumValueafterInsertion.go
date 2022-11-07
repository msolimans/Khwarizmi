package maximumvalueafterinsertion

import (
	"strconv"
	"strings"
)
func maxValue(n string, x int) string {
    var sb strings.Builder 

    inserted := false 
    isNegative := false 
    if n[0] == '-' {
        isNegative = true 
    }

    for i := 0;i<len(n);i++ {
        if isNegative && i == 0 {
            sb.WriteRune('-')
            continue 
        }

        //if its negative make sure to put the digit right before the digits that's greater than x 
		if isNegative && x < int(n[i] - '0') && !inserted  { 
            sb.WriteString(strconv.Itoa(x))
            sb.WriteByte(n[i])
            inserted = true  
        } else if !isNegative && x > int(n[i] - '0') && !inserted {//if positive, make sure to put the digit right after the digits that's greater than x 
            sb.WriteString(strconv.Itoa(x))
            sb.WriteByte(n[i])
            inserted = true 
        } else {
            sb.WriteByte(n[i])
        }
    }

	//if not inserted at all, add it to the end 
    if !inserted {
        sb.WriteString(strconv.Itoa(x))
    }

    return sb.String()
}

//thoughts 
// if positive => put it right after the digit that's greater than x 
//99 and 8 = 9 > 8 => 9 > 8 => 998
// -55889 7 
// -99875 7 
// if negative, put it right before the digit that's greater than x 
//-55 and 9 
// 9 > 5 => 9 > 5 => 559
//-45 and 3 => 3 > 4 (no) insert before 4 => -345



