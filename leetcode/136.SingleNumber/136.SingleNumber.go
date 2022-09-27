package singlenumber

//xor should remove dups, the left number will be the result
func singleNumber(nums []int) int {
	res :=0 
	for _,n := range nums{
		res ^= n 
	}

	return res 
}