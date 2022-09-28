package removingstarsfromastring

import (
	"container/list"
	"strings"
)

func removeStars(s string) string {
	//use stack to remove the last-in char => first out :) 
	//* always comes up after characters (I have tested "***s" and it gave me invalid test case in leetcode)
	//in constraints, we see the lenght is greater than 1 so we don't have t worry about empty string inputs 
	
	//simple 
	stack := list.New()
	for _, c := range s {
		// fmt.Println(reflect.TypeOf(c))
		if c == '*' {
			stack.Remove(stack.Front())
		} else {
			stack.PushFront(c)
		}
	}

	var sb strings.Builder
	for stack.Len() > 0 {
		back := stack.Back()
		sb.WriteRune(back.Value.(int32))
		stack.Remove(back)
	}

	return sb.String()
}