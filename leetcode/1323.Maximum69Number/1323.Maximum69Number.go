package maximum69number

import "container/list"

func maximum69Number (num int) int {
    // 9669 \ 6699 -> 9699 
    stack := list.New()

    for num > 0 {
        stack.PushFront(num % 10)
        num /= 10 
    }

    res := 0 
    pass := false 
    for stack.Len() > 0 {
        f := stack.Front()
        digit := f.Value.(int)
        stack.Remove(f)

        if digit == 6 && !pass {
            digit = 9 
            pass = true 
        }

        res = res*10 + digit
    }

    return res 
}