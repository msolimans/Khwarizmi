package countprefixesofagivenstring

func countPrefixes(words []string, s string) int {
    mp := map[string]byte{}
    
    for i :=1;i<=len(s);i++ {
        mp[s[0:i]] = byte(0)
    }
    
    c := 0 
    for _,w := range words {
        if _,exists := mp[w]; exists {
           c++  
        }
    }
    
    
    return c 
    
}

