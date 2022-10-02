package merge

func sort(arr []int) []int {

	sortInPlace(arr,0,len(arr)-1)
	return arr 
	
}


//with start and end 

func sortInPlace(arr []int, s, e int) {
	if s >= e {
		return
	}


	mid := s + (e - s) / 2

	sortInPlace(arr, s, mid)
	sortInPlace(arr, mid+1, e)
	merge(arr,s, mid, e)
	 
}

//merge 2 sorted arrs (in place)
func merge(arr []int,start,mid,end int) {
	// [1,8,9,2,3,4]
	//0,2 and 3,5

	// we check both sides, if left is less than right, move on otherwise
	// shift to the left (from current end to current right and fill in current left with min val we got from right side)	
	if arr[mid] < arr[mid+1] {
		return //already sorted 1,2,   3,4
	}

	// s1 is s 
	start2:=mid+1
	for start <= mid && start2 <= end {
		//we check 
		if arr[start] <= arr[start2] {
			start++ //move start 
		} else {
			//shift to the right and update poitners 

			val  := arr[start2]
			indx := start2 

			for indx != start  { //move until current's next node 
				arr[indx] = arr[indx - 1]
				indx--
			}
			arr[start] = val 

			//shift all pointers to right  (don't change end point at all )
			start++
			mid++ 
			start2++ 
		}


	}
	

}

