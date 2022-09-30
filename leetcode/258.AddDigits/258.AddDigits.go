package adddigits

func addDigits(num int) int {
    
    for num >= 10 {
        num = add(num)        
    }
    
    return num 
}

func add(num int) int {
    sum := 0
    for num > 0 {
        sum += num % 10 
        num /= 10 
    }
    
    return sum 
}