package checkifstringisaprefixofarray

func isPrefixString(s string, words []string) bool {
    i := 0 
    j := 0 
    k := 0 
    
    for i < len(s) {
        //at anytime if j exceeded len of words and we still have chars in s - we are done 
        if j > len(words) -1 {
            return false 
        }
        
        if s[i] != words[j][k] {
            return false 
        }
        
        //move pointers 
        i++
        if k == len(words[j]) - 1 {
            k = 0 //reset k 
            j++ //move to next word
        } else {
            k++
        }
    }
    
    return k == 0
}
