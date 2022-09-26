package excelsheetcolumntitle

import "strings"

func convertToTitle(columnNumber int) string {
	
	//convert from number to string 
	// keep dividing by 26 
	// remainder + 64 => convert to char 
	// othereise it's Z => division result go to nxt round 
	var sb strings.Builder
	for columnNumber != 0 {
		if columnNumber % 26 == 0 {//it's z
			sb.WriteString("Z")
			columnNumber = (columnNumber/26) - 1
		} else {//remainder (calc which char)
			whichChar := rune((columnNumber % 26 ) + 64)
			sb.WriteRune(whichChar)
			columnNumber /= 26
		}
	}
	return reverse(sb.String())

    
}

func reverse(s string) string {
	var sb strings.Builder

	for i:=len(s)-1;i>=0;i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}