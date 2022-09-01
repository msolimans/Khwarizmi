package datastructure

type Queue []string 

func (q *Queue) Enqueue(input string ) {
	*q = append(*q, input)
}

func (q *Queue) Dequeue() string { 
	res := (*q)[0]
	*q = (*q)[1:]
	return res 
}

func (q *Queue) IsEmpty () bool {
	return len(*q) == 0 
}

func (q *Queue) Len() int  { 
	return len(*q) 
}

/////////
// Above Implementation can have memory leaks. We may add 3 items in array and pop only 2 leaving the first one resulting in memory leaks
// Its better to use LinkedList
