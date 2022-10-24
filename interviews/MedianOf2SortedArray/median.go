package median

import "fmt"

func calcMedian(arr1, arr2 []int) float64 {

	n1,n2 := len(arr1) - 1, len(arr2) - 1
	m1,m2 := n1 / 2 , n2 / 2 
	//3 cases 
	//first array comes before 2nd 
	//0 1 2  3  4  5  6 
	//1 2 3  4  5  6  7
	//8 9 10 11 12 13 14

	//7+8 / 2 

	if arr1[n1] < arr2[0] { //arr1 before arr2 
		return float64(arr1[n1] + arr2[0]) / 2 
	}
	
	//last elem from first and first from 2nd 
	//first array comes after 2nd 
	
	//0 1 2  3  4  5  6 
	//8 9 10 11 12 13 14
	//1 2 3  4  5  6  7 

	//7+8 / 2 

	if arr2[n2] < arr1[0] { //arr1 after arr2 
		return float64(arr1[0] + arr2[n2] ) / 2 
	}

	//in between (overlap) => find median of both 
	if (n1 + 1) % 2 > 0 {
		if arr1[m1] < arr2[m2] {
			return  calcSumOdd(arr1, arr2, m1, m2) / 2
		}  
		return calcSumOdd(arr2, arr1, m2, m1) / 2
	}

	//even
	if arr1[m1] > arr2[m2] {
		return  calcSumEven(arr1, arr2, m1, m2) / 2
	}  
	return calcSumEven(arr2, arr1, m2, m1) / 2
}

//naive solution (using it for unit testing)
func calcMedian2(arr1, arr2 []int) float64 {

	merged := make([]int, len(arr1) + len(arr2) ) 
	c, i,j := 0,0,0
	 
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged[c] = arr1[i]
			i++
		} else {
			merged[c] = arr2[j]
			j++ 
		}

		c++ 
	}

	for i < len(arr1) {
		merged[c] = arr1[i]
		i++ 
		c++ 
	}


	for j <len(arr2) {
		merged[c] = arr2[j]
		j++ 
		c++ 
	}

	//calc 
	m := (len(merged) - 1) / 2
	fmt.Println(merged[m] , merged[m+1])
	return float64(merged[m] + merged[m+1]) / 2 
}

func calcSumEven(arr1 []int, arr2 []int, m1 int, m2 int) float64 {
	//assume arr1[m1] is bigger than arr2[m2]
	//choose which one is bigger 
	//then check previous bigger vs next smaller => pick the bigger 
	res := 0 
	
	res += arr1[m1] //m1 is bigger 
	if m1 == 0 {//no previous 
		return float64(arr1[m1] + arr2[m2 + 1]) 
	}

	if arr1[m1-1] > arr2[m2 + 1] { //previous to bigger is greater than next to smaller we choose previous 2 4 6 10 and 8 9 10 12 
		return float64(arr1[m1] + arr1[m1-1])
	} else {
		return float64(arr1[m1] + arr2[m2 + 1])
	}
}
func calcSumOdd(arr1 []int, arr2 []int, m1 int, m2 int) float64 {
	
	//i don't think we will have this case since ln 14 and 27 will handle this 
	if m1 == len(arr1)  - 1 { //[1] - [2] OR [4] - [3]
		return float64(arr1[m1] + arr2[m2]) //case where there's only 1 elem in the array 
	}  
		
	res := arr1[m1]
	//who's less (next to arr1[m1] or arr2[m2])
	if arr1[m1+1] < arr2[m2] {
		res += arr1[m1+1]
	} else {
		res += arr2[m2]
	}
	
	return float64(res)
}