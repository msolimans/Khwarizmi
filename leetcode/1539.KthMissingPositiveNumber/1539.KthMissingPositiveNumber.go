package kthmissingpositivenumber

func findKthPositive(arr []int, k int) int {
	iMiss := 0 //how many misses (when its equal to k return iNo)
	iNo := 1 //index of number 
	iArr := 0 //the index of array 

	for iArr < len(arr) {
		if arr[iArr] != iNo {
			iMiss++

			if iMiss == k {
				return iNo
			}
			iNo++
			//dont move iArr (only when we encounter equality)
		} else {
			iNo++ //1 2 3 4 5 6 
			iArr++ //1 2 3 5
		}

	}

	return (k - iMiss) + arr[len(arr) - 1]
    
}
