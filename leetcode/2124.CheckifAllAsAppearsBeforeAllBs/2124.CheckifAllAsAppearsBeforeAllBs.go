package checkifallasappearsbeforeallbs


func checkString(s string) bool {
    
	//we have not visited b 
    b := false 
    for _,c := range s {
		//once visited b, mark b as true  
        if c == 'b' {
            b = true 
        } else if c == 'a' && b { //if we visit a after b being marked as `true`, return false 
            return false 
        }
    }
    
	//return true 
    return true 
}