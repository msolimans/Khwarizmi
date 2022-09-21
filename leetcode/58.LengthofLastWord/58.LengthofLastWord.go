package lengthoflastword

func lengthOfLastWord(s string) int {
    //spaces 
    i := len(s) - 1
    for ;i>=0;i-- {
        if s[i] != ' '{
            break
        }
    }
    
    en := i
    for ;i>=0;i-- {
        if s[i] == ' ' {
            break
        }
    }
    st := i+1
    
    return en - st + 1
    
}