package findkclosestelements

import (
	"math"
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {
    
    //similar to find closest match - use binary search (o(logn)
    indx := findClosestIndex(arr, x, 0, len(arr) - 1)
	
	if indx == 0 {
		return expand(arr, indx,indx+1, k, x)
	} else if indx == len(arr) - 1 {
		return expand(arr, indx-1,indx, k, x)
	} else {
		return expand(arr, indx-1,indx, k, x)
	}
}

func findClosestIndex(arr []int, x, l, r int) int {
    mid := l
    
	for l <= r {
        
        mid = (l + r) / 2

        if arr[mid] == x {
            return mid 
        }  else if arr[mid] < x {
            l = mid + 1
        } else if arr[mid] > x {
            r = mid - 1
        }
    }
    
	if x > arr[mid] {
		return mid + 1
	} else {
		return mid 
	}
}

func expand(arr []int, i,j, k, x int)  []int {
	 

	res := []int{}
	for k > 0 {
		if i >= 0 && j <= len(arr) - 1 {
			if math.Abs(float64(arr[i]) - float64(x)) <=  math.Abs(float64(arr[j]) - float64 (x)) { 
				res = append(res, arr[i])
				if i == j {//first case (move both)
					j++ 
				}
				i-- 
				k--
			} else {
				res = append(res, arr[j])
				j++
				k--
			}
		} else if i >= 0 {
			res = append(res, arr[i])
			i--
			k--
		} else if j <= len(arr) - 1 {
			res = append(res, arr[j])
			j++
			k--
		} else {
			break //out of boundary 
		}
		
	}

	sort.Ints(res)

	return res
}