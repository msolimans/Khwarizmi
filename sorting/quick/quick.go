package quick

func sort(arr []int) {
	quickSort(arr,0,len(arr)-1)
}

func quickSort(arr []int, s,e int) {

	if s>= e{
		return
	}

	pivot :=  partition(arr, s, e)
	quickSort(arr, s, pivot-1)
	quickSort(arr, pivot+1, e)

}

func partition(arr []int, s,e int) int {
	//pivot is the at index end 
	//partition array so that every element before pivot are less than the pivot and every elems after pivot are greater than pivot 

	pivot := arr[e]

	//j points to the last element in the loop 
	//i points to the begining of bigger elems in the array (where we last put pivot at that index)
	i, j := s, s 

	for j = s; j<=e-1;j++ { //pivot is excluded
		if arr[j] <= pivot {
			swap(arr, i,j)
			i++
		} 
	}
	swap(arr, i, e) //swap last elem to put pivot in the middle 
	return i 
}

func swap(arr []int, i,j int) {
	t := arr[i] 
	arr[i] = arr[j] 
	arr[j] = t 
}