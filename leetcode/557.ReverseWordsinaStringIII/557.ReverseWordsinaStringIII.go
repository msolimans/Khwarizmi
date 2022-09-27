package reversewordsinastringiii

import "strings"

func reverseWords(s string) string {

	strs := strings.Split(s, " ")
	var sb strings.Builder
	
	for _, str := range strs {
		sb.WriteString(reverseWord(str))
		sb.WriteString(" ")
	}

	res := sb.String()
	return res[0:len(res)-1]
}

func reverseWord(s string) string {
	chs := make([]byte, len(s))
	for i := 0;i<len(s);i++ {
		chs[i] = s[i]
	}

	i, j := 0, len(chs)-1

	for i < j {
		t := chs[i]
		chs[i] = chs[j]
		chs[j] = t 
		i++
		j--
	}

	return string(chs)
}
