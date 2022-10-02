package selection

//worst case scenario is O(n^2) in which we have sorted array in descending order and we want to sort it in an ascending order.
func sort(arr []int) []int {
	//for every elem, we make sure all elems coming after are greater than current element 
	for i :=0;i<len(arr);i++{
		for j:=i+1;j<len(arr);j++ {
			if arr[i] > arr[j] {
				//swap 
				swap(arr, i,j)
			}

		}

	}

	return arr 

}

func swap(arr []int, i, j int) {
	t := arr[i] 
	arr[i] = arr[j] 
	arr[j] = t 
}