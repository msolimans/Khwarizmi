package minimumsumoffourdigitnumberaftersplittingdigits

import "sort" 

func minimumSum(num int) int {
    arr := []int{}
    
	//get the lowest number in the left and the highest number in the right
	//in other words, 2923 => 2,2,3,9 => 2 goes with 9 and 2 goes with 3 => 29+23=52 

	//convert number to array 
    for num > 0 {
        arr = append(arr, num % 10) 
        num /= 10 
    }
    //sort array 
    sort.Ints(arr)
    
	return merge(arr[0], arr[3]) + merge(arr[1], arr[2])
}


func merge(i,j int) int {
    return i*10+j
}

// 9842

// 29
// 48

