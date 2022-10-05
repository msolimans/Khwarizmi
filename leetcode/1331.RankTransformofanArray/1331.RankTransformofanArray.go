package ranktransformofanarray

import "sort"

func arrayRankTransform(arr []int) []int {
    if len(arr) == 0 {
        return []int{}
    }
	
	//clone array 
    sorted := make([]int, len(arr) )
    for i :=0; i<len(arr);i++ {
        sorted[i] = arr[i]
    }
    
	//sort array 
    sort.Ints(sorted)
    
	//map that stores each element with its rank  
	mp := map[int]int{} 
    count := 1 
    current := sorted[0]
    mp[current] = count 
    
	//calc rank - example given below: 
	//37,12,28,9,100,56,80,5,12
	//====
	//5,9,12,12,28,37,56,80,100
	//1,2,3,3,4,5,6,7,8
    for i := 1; i < len(sorted); i++ {
        if current == sorted[i] {
            continue 
        } 
        
        count++
        current = sorted[i]
        mp[current] = count
    }
    
	//store array's rank and return it 
    for i := 0;i < len(arr);i++ {
        arr[i] = mp[arr[i]]
    }
    
    return arr 
}

