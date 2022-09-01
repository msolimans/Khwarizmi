package LongestCommonPrefix

func longestCommonPrefix(strs []string) string {
	matched := strs[0]
	for i :=0;i<len(strs);i++ {
		matched = findMatch(matched, strs[i], 0)
	}

	return matched 
    
}

func findMatch(s1 string,s2 string, i int) string {
	if i >= len(s1) || i >= len(s2) {
		return s1[0:i]
	}

	if s1[i] != s2[i] {
		return s1[0:i]
	}
	i+=1
	return findMatch(s1, s2, i)
}