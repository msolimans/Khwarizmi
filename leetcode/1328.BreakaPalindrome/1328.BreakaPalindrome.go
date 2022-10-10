package breakapalindrome

func breakPalindrome(palindrome string) string {

	if len(palindrome) == 0 || len(palindrome) == 1 {
		return ""
	}
    
	i := 0
	for ;i<len(palindrome)/2 ;i++{
		if int(palindrome[i]) - 'a' > 0 { //aaaaa 
			return palindrome[0:i] + "a" + palindrome[i+1:]
		} 
    }

	return palindrome[0:len(palindrome)-1] + "b"

	
}