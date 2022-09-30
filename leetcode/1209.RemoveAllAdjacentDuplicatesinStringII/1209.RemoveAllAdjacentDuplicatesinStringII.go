package removealladjacentduplicatesinstringii

import (
	"container/list"
	"strings"
)

type StackItem struct {
	Val rune 
	Count int 
}

func removeDuplicates(s string, k int) string {
	//we use stack and then we store how many stacked up at each point 
	stack := list.New()
	for _,c := range s {
		
		if stack.Len() > 0 {
			sf := stack.Front()
			item := sf.Value.(*StackItem)
			if item.Val == c && item.Count == k - 1 {
				stack.Remove(sf)
			} else if item.Val == c {
				//keep it and add 1 to count 
				item.Count += 1
			} else {//not equal 
				stack.PushFront(&StackItem {
					Val: c,
					Count: 1,
				}) 
			}
		} else {
			stack.PushFront(&StackItem {
				Val: c,
				Count: 1,
			})
		}
	}

	
	//read from the back of the list to write up string normally 
	var sb strings.Builder
	for stack.Len() > 0 {
		back := stack.Back()
		//we may have dups -> like k = 3 and count as 2 in stack for char 'a' so we output "aa" (list them all)
		for i :=0 ;i<back.Value.(*StackItem).Count;i++ {
			sb.WriteRune(back.Value.(*StackItem).Val)
		}
		stack.Remove(back)
	}

	return sb.String()
}

