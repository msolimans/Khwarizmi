package makethestringgreat

import "container/list"

func makeGood(s string) string {
    //use stack to remove two adjacent 
    //every time you encounter a char look at stack (peek) in case it is similar then remove otherwise insert it 
    stack := list.New()

    for i := 0;i<len(s);i++ {
        if canInsert(stack, s[i]) {
           stack.PushFront(s[i]) 
        } else {
            //remove in case we have any elems in stack 
            if stack.Len() > 0 {
                stack.Remove(stack.Front())
            }
        }
    }

    //reversed result in stack 
    res := make([]byte, stack.Len())
    i := len(res) - 1  
    for stack.Len() > 0 {
        f := stack.Front()
        res[i] = f.Value.(byte)
        stack.Remove(f) 
        i-- 
    }

    return string(res) 
}

// func reverse(r []rune) string {
//     i, j := 0, len(r) - 1 
//     for i < j {
//         r[i], r[j] = r[j], r[i]
//         i++
//         j--
//     }
//     return string(r)
// }

func isCapital(i byte) bool  {
    if i - 'A' >= 0 && i - 'A' <= 26 { //capital 
     return true 
    }
    return false 
}

func canInsert(stack *list.List, i byte) bool {
    if stack.Len() == 0  {
        return true 
    }

    peek := stack.Front().Value.(byte) 
    iIsCapital := isCapital(i)
    peekIsCapital := isCapital(peek) 

    r1 := i - 'A'
    if !iIsCapital { //capital 
        r1 = i - 'a'
    }
    
    r2 := peek - 'A'
    if !peekIsCapital { //capital 
        r2 = peek - 'a'
    }

    if r1 == r2 { //both are equal 
        if iIsCapital && peekIsCapital { //both are capitals 
            return true
        }

        if !iIsCapital && !peekIsCapital {//both are small 
            return true 
        }

        //  one of them are capital and other is small
        return false //remove in such case from stack and ignore current elem 
    }

    return true 
}