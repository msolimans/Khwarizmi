package fizzbuzz

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	res := make([]string, n)

	for i:=0;i<n;i++ {
		no := i+1
		if no % 3 == 0 && no % 5  == 0 {
			res[i] = "FizzBuzz"
		} else if no % 3 == 0 {
			res[i] = "Fizz"
		} else if no % 5 == 0 {
			res[i] ="Buzz"
		} else {
			res[i] = strconv.Itoa(no)
		}

	}

	return res 
    
}