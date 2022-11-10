package printwordsvertically

import "strings" 

func printVertically(s string) []string {
    arr := strings.Split(s, " ")
    
    j := 0 
    passed := true 
    res := []string{}
    //loop thru strs and 
    //in each round contstruct a string and add it 
    for passed {
        passed = false 
        round := []byte{} 
        for i := 0;i<len(arr);i++ {
            if j > len(arr[i]) - 1 {//we exceeded length of string so add empty space 
                round = append(round, ' ')
            } else {
                round = append(round, arr[i][j]) //we add the char at this position from arr[i]
                passed = true //we mark we passed this round indicating we have at least one char
            }
        }
        if !passed { //no chars at all (round should be full of spaces so ignore it)
            break
        } else {
            fr := string(round)
            if fr[len(fr) - 1] == ' ' {
                fr = strings.TrimRight(fr, " ")
            }
            res = append(res, fr)
        }
        j++
        
    }

    return res 
     

}