package replaceelementswithgreatestelementonrightside

// 0   1 2 3 4 5
// 17 18 5 4 6 1

// 5 - -1
// 4 - 1
// 3 - 6
// 2 -

func replaceElements(arr []int) []int {
	// O(n) time complexity with O(1) space 
    //backward moves are better to get save max and update current index based on max 
	max := -1 
    for i := len(arr) - 1;i>=0 ;i-- {
        t := arr[i]
        arr[i] = max 
        if t > max {
            max = t 
        }
    }
    
    return arr 
}
