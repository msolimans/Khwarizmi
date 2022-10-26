package medianoftwosortedarrays


func findMedianSortedArrays(arr1, arr2 []int) float64 {
 
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
	//fmt.Println(merged[m] , merged[m+1])
    if len(merged) % 2 == 0 {
        return float64(merged[m] + merged[m+1]) / 2 
    }
    return float64(merged[m])
}
 