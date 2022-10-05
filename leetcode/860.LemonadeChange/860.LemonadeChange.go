package lemonadechange

func lemonadeChange(bills []int) bool {
    //5 10 20 
    fives := 0 
    tens := 0 
    twenties := 0 
    
    
    for _,b := range bills {
        if b == 5 {
            fives++
        }  else if b == 10 && fives > 0 {
            tens++ //add ten
            fives-- //owe five 
        }  else if b == 20 && fives > 0 && tens > 0 {
            twenties++ //add 20s (not sure if we need it)
            tens--
            fives--
        } else if b == 20 && fives >= 3 {
            twenties++
            fives -= 3
        } else {
            return false 
        }
    }
    
    return true 
}




//miss understand question - the following solution is based off cumulative/total changes available 
func lemonadeChange2(bills []int) bool {
    //5 10 20 
    fives := 0 
    tens := 0 
    twenties := 0 
    
    //owes fives or tens 
    minus_fives := 0 
    minus_tens := 0 
    
    for _,b := range bills {
        if b == 5 {
            fives++
        }
        if b == 10 {
            tens++ //add ten
            minus_fives++ //owe five 
        }
        if b == 20 {
            twenties++ //add 20s (not sure if we need it)
            minus_tens++ // this can be 2 fives
            minus_fives++
        }
    }
    
    //mandatory to have fives here 
    if fives < minus_fives {
        return false 
    } else {
        fives -= minus_fives
    }
    
    //whatever left here can come from tens and fives 
    if (fives / 2) + tens < minus_tens {
        return false 
    }
    
    return true 
}