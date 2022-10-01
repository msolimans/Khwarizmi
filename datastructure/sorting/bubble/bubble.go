package bubble

func sort(arr []int) []int {
	// compare 2 elems in every loop until we don't do any swaps 

	
	for i :=0;i<len(arr);i++ {
		//assume array is sorted 
		sorted := true 

		for i :=0;i<len(arr)- 1;i++ {
			if arr[i] > arr[i+1] {
				//swap (if we hit one swap it means it is not sorted yet, eval in future rounds) 
				swap(arr, i , i+1)
				sorted = false
			}
		}

		//once sorted, break 
		if sorted {
			break
		}
	}

	return arr 

}

func swap (arr []int, i,j int) {
	t := arr[i] 
	arr[i] = arr[j] 
	arr[j] = t 
}
