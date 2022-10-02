package sumofdigitsofstringafterconvert

import (
	"strconv"
	"strings"
)

func getLucky(s string, k int) int {
	//convert chars to ints (stored in array of string)
	//this is needed for big strings like "hvmhoasabaymnmsdzzzzhvmhoasabaymnmsdzzzzhvmhoasabaymnmsdzzzzhvmhoasabaymnmsdzzzz"
	//otherwise we get int overflow 
	str := convertToNumbersAsStr(s) 
	
	res := 0
	//sums all numbers in string (and return sum as numbers ir integer)
	res = sum(str)
    
    //sums numbers (k - 1 times) 
	for i := 1;i<k;i++ {
		res = sumInts(res)
	}

	return res 
}

//sum ints 
func sumInts(n int) int {
	sum := 0 
	for n > 0 {
		sum += n % 10 
		n /=10
	}
	return  sum 
}

//sums up all numbers together e.g. 9999 => 9 + 9 + 9 + 9 
func sum(s string) int {
	sum := 0

	for i:=0;i<len(s);i++ {
		v, _ := strconv.Atoi(string(s[i]))

		sum += v 
	}

	return sum 
}

//converts chars into string array of numbers 
func convertToNumbersAsStr(s string) string {

	var sb strings.Builder
	for i := 0;i<len(s);i++ {
		n := (int(s[i]) % 97) + 1 //1 2 3 ...26 

		sb.WriteString(strconv.Itoa(n))

	}

	return sb.String()
}

