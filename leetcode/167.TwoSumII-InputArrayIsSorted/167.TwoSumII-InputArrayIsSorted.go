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