package minimuminsertionstobalanceaparenthesesstring

import "container/list"

func minInsertions(s string) int {
    stack := list.New() 
    
    count := 0 
    for i :=0;i<len(s);i++ {
        if s[i] == '(' {
            stack.PushFront(s[i])
        } else { //it is ) since s has only ( and ) 
            //now we pop and move one step more 
            f := stack.Front()
            var ch uint8 = ' '
			
			if f != nil {
				ch = f.Value.(uint8)
			}
            
			if ch == ' ' {
				count++ //for opening brace 
                
                //we check if this current and next one is ) 
                if i == len(s) - 1 {
                    count++  //2nd closing  
                } else if s[i+1] == ')' {
                    i+=1 
                } else {
					count++ //2nd closing 
				}
			} else { //if fv.(uint) == '(' 
                //we check if this current and next one is ) 
                if i == len(s) - 1 {
                    count++ 
                } else  if s[i+1] == ')' {
                    i+=1 
                } else {
					count++
				}
                //pop 
                 stack.Remove(f)
            }  
            
        }
    }

	for stack.Len() > 0 {
		stack.Remove(stack.Front())
		count+=2
	}
    
    return count 
}