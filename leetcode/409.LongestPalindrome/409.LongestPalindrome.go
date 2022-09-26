package main


func longestPalindrome(s string) int {
    chars := make([]int, 52) 
    
    for _, c := range s {
    	chars[c % 52] += 1
    }

	hasOdd := false 
	count := 0 
	for i :=0;i<len(chars);i++ {
		if chars[i] % 2 == 0 {
			count += chars[i]
		} else {
			hasOdd = true 
			//in case it's 3, count 2 and ignore the rest 
			count += (chars[i] / 2) * 2
			
		}

	}
    
	if hasOdd {
		count++ 
	}

	return count 
    
}

