package countingwordswithagivenprefix

func prefixCount(words []string, pref string) int {
    c :=0
    for _, s := range words {
        if compare(s,pref) {
         c++   
        }
    }
    return c 
}

func compare(s, pref string) bool {
    if len(s) < len(pref) {
        return false 
    }
    
    i := 0
    
    for i< len(pref) {
      
        if s[i] != pref[i] {
            return false 
        }
        i++
    }
    
    return true 
}