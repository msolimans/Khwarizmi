package countintegerswithevendigitsum

//fill in array with "even" digitsums numbers as helper
var arr = make([]bool, 1001)

func countEven(num int) int {
    if len(arr) == 0 || arr[2] == false {
        for i := 0;i<len(arr);i++ {
            arr[i] = isDigitSumEven(i)
        }
    }
    
    c := 0 
    for i := num;i>=1;i-- {
        if arr[i] {
            c++
        }
    }
    
    
    return c 
}


func isDigitSumEven(num int) bool {
    c := 0 
    
    for num > 0 {
        c += num % 10 
        num /= 10 
    }
    
    return c % 2 == 0 
}

