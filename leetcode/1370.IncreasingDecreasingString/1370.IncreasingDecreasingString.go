package increasingdecreasingstring

import "strings" 

func sortString(s string) string {
    
    //97 
    var sb strings.Builder 
    buckets := make([]int,28)
    
    for _,v := range s {
        buckets[(v - 97) + 1] += 1
    }
        
    done := false 
    i := 0
    dir := 1
    
    for {
        if i == len(buckets) - 1 {
            if done {
                break
            } else {
                done = true  //suppose we will be done by this round 
            }
            dir = -1
        }
        
        if i == 0 {
            if done {
                break
            } else {
                done = true  //suppose we will be done by this round 
            }
            dir = 1
        }
        
        if buckets[i] > 0 {
            sb.WriteRune(rune(i + 96))
            buckets[i] = buckets[i] - 1
            done = false 
        }
        
        i += dir 
    }
    
    return sb.String()
}