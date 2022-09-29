package makearrayzerobysubtractingequalamounts

func minimumOperations(nums []int) int {
		//count how many distinct elems in the array - should give us how many rounds we can use to make it zero 
		arr := make([]int, 101)
		
		for _, n := range nums {
			if arr[n] > 0 {
				continue 
			}
			arr[n] = 1
		}
		
		count := 0 
		for i := 1;i<len(arr);i++ {
			if arr[i] == 1 {
				count++
			}
		}
		
		return count 
	
	}
	
	
	
	