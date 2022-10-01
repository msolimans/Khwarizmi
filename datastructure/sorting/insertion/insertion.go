package insertion

func sort( arr []int) []int {
	//in every loop, we put (insert) the element at its correct position)
	//start at position 1 
	for i :=1;i<len(arr);i++ {
		//left elem is less than current 
		if arr[i-1] <= arr[i]  {
			continue
		}

		//insert current element where it should be 
		x := i 
		y := i - 1
		for x>0 && arr[y]>arr[x] {
			//swap 
			swap(arr, x, y)
			x--
			y-- 
		}

	}

	return arr 
}


func swap(arr []int, i,j int) {
	t := arr[i] 
	arr[i] = arr[j] 
	arr[j] = t 
}
