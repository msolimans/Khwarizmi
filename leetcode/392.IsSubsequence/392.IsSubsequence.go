package issubsequence

func isSubsequence(s string, t string) bool {
    is := 0
	it := 0 
	if len(t) < len(s){
		return false
	}

	for it < len(t) && is < len(s) {
		if s[is] != t[it] {
			it++
		} else {
			it++
			is++
		}
	}

	return is == len(s) 
}