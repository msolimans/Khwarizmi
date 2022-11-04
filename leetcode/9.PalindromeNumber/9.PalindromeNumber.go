package palindromenumber

func isPalindrome(x int) bool {
	//reverse number and check if they are both equal. if so, return true otherwise return false
    reversed := 0 
    original := x 

    for original > 0 {
		//bring from right most of number 
        rightMost := original % 10
		//put it in begining of reversed (make a room for right most by multiplying by 10)
        reversed = reversed * 10 + rightMost
        original /= 10 
    }

    return reversed == x 

}