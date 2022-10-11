package uniqueemailaddresses

import (
	"strings"
)

func numUniqueEmails(emails []string) int {
    mp := map[string]bool{} 
    
    for _,e := range emails {
        e = filter(e)
        if _,exists := mp[e]; !exists {
            mp[e] = true 
        }
    }    
    
    
    return len(mp)
}

//split strings around @ 
//ignore anything after + 
//replace all . with ""
//combine whatever we need
func filter(email string) string {
    earr := strings.Split(email, "@")
    
    locals := strings.Split(earr[0], "+")

	return strings.ReplaceAll(locals[0], ".", "") + "@" + earr[1]
}
