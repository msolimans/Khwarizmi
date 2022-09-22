package isomorphicstrings

import (
	"strconv"
	"strings"
)

func isIsomorphic(s string, t string) bool {
    
    if len(s) != len(t) {
        return false 
    }
    
    mp := map[byte]byte{}
    mpv := map[byte]byte{}
    
    for i := 0;i<len(s);i++ {
        //as a key 
        if val, exist := mp[s[i]]; exist {
            if val != t[i] {
                return false 
            }
        } else {
            //it doesn't exist as a key (good!) now we make sure it doesn't exist in values (we keep track of values in different map) 
            if _, exists := mpv[t[i]]; exists {
                return false 
            }
        }
        
        //add chars from both s and t 
        mp[s[i]] = t[i]
        mpv[t[i]] = s[i]
    }
    
    
    return true 
}
  
func isIsomorphic2(s string, t string) bool {
	return transform(s) == transform(t)
}


//transform string to integers based on first occurence transformation 
//addat
//01104
//errea
//01104 (true)
//errsa
//01134 (false)

//paper 01034
//titee 01033
func transform(s string) string {
	var sb strings.Builder
	mp := map[byte]int{}

	for i:=0;i<len(s);i++ {
		if v, exists := mp[s[i]]; exists {
			sb.WriteString(strconv.Itoa(v))
		} else {
			mp[s[i]] = i 
			sb.WriteString(strconv.Itoa(i))	
		}
	}

	return sb.String()
}