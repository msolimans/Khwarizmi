package calculatedigitsumofastring

import (
	"strconv"
	"strings"
)

func digitSum(s string, k int) string {
	
	if len(s) <= k {
		return s 
	}

	var sb strings.Builder
	i:=0
	for  ;i<len(s)-k;i+=k {
		sb.WriteString(sum((s[i:i+k])))
	}

	if i < len(s) {
		sb.WriteString(sum(s[i:]))
	}



	return digitSum(sb.String(), k)
	
	
}

func sum(s string) string {
	sum := 0 
	for i:=0;i<len(s);i++ {
		v,_:=strconv.Atoi(string(s[i]))
		
		sum += v 
	}

	return strconv.Itoa(sum) 
}