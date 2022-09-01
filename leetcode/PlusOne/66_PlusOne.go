package plusone

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1;i>=-1;i--{ 
		if i < 0 && carry > 0 {
			finalRes := make([]int, len(digits) + 1)
			finalRes[0] = carry
			//copy all from digits (since it was updated )
			for j:=0;j<len(digits);j++{
				finalRes[j+1] = digits[j]
			}

			return finalRes
		} else if carry > 0 {
			plus := digits[i] + carry
			digits[i] = plus % 10 
			carry = plus / 10
		} else {
			break
		}
	}
    
	return digits
}
