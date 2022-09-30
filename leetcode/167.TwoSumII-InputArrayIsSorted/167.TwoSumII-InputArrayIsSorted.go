package twosumiiinputarrayissorted

func twoSum(numbers []int, target int) []int {
    //since we have array sorted, 2 pointers where we move such pointers according to the target 
    i,j := 0,len(numbers)-1
    
    sum := 0
    for i < j {
        sum = numbers[i] + numbers[j]
        if sum == target {
            return []int{i+1,j+1}
        }else if sum > target {
            //move j to the left (to decrease the total sum)
            j--
        }  else {//less than 
            i++
        }
    }
    
    return []int{}
    
    
}



//second solution (using hashmaps) 
func twoSum2(numbers []int, target int) []int {
    mp := map[int]int{} 
    
    //save into map (pair of numbers and their indexes)
    for i:=0;i<len(numbers);i++ {
        mp[numbers[i]] = i 
    }
    
    //now we look at each one and compare it with the target 
    for i := 0;i<len(numbers);i++ {
        goal := target - numbers[i] 
        if _, exists := mp[goal]; exists {
            if i < mp[goal] {
                return []int{i+1, mp[goal]+1}
            } else {
                return []int{mp[goal]+1, i+1}
            }
        }
    }
    
    return []int{}
}
