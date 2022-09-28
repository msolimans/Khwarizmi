package removealladjacentduplicatesinstring

import (
	"container/list"
	"strings"
)

func removeDuplicates(s string) string {
	//use stack to push chars (in case we see dups in top of stack both chars should cancel each other)
	stack := list.New()
	
	
	for _, c := range s {

		if stack.Len() == 0  || stack.Front().Value.(rune) != c {
			//push it and move forward 
			stack.PushFront(c)
		} else { //similar
			//pop from front of linked List 
			stack.Remove(stack.Front())
		}
	}

	var sb strings.Builder
	for stack.Len() != 0 {
		//get from back of linked List (now it should be dealt with as a queue)
		sb.WriteRune(stack.Back().Value.(rune))
		stack.Remove(stack.Back())
	}

	return sb.String()

}
