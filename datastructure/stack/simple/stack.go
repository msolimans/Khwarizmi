package simple

//first in, last out
type Stack []string

func (s  *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Len() int { 
	return len(*s) 
}
//append 
func (s *Stack) Push(input string) {
	*s = append(*s, input) 
}

//pop out from last then shrink in array of strings 
func (s *Stack) Pop() string { 
	index := len(*s) - 1
	res := (*s)[index]
	*s = (*s)[:index]
	return res 
}
