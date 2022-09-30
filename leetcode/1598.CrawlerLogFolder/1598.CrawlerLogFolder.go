package crawlerlogfolder

import "regexp"

func minOperations(logs []string) int {
    //we add to the count only in case we move deeper inside folder/dir structure 
	//we subtract from count everytime we go one level up (make sure it doesn't go before 0)
	count := 0
    min := regexp.MustCompile(`\.\./`)
    plus := regexp.MustCompile(`[0-9a-z]+/`)
    
    for _, l := range logs {
        if min.Match([]byte(l)) {
            if count > 0 {
                count--
            }
        } else if plus.Match([]byte(l)) {
            count++
        }
    }
    
    return count 
    
}