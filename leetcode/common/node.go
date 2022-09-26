package common

import (
	"strconv"
	"strings"
)

type Node struct {
	Val      int
	Children []*Node
}	



//serialize it here 
func (n *Node) ToString() string {	
	var sb = &strings.Builder{}
	iterativeString(n, sb)
	return sb.String()
}

//looks ugly but performant than concatenating immutable strings 
func  iterativeString(n *Node, res *strings.Builder) {
	if n == nil {
		return 
	}

	res.WriteString(strconv.Itoa(n.Val))
	for _, c:= range n.Children {
		iterativeString(c, res)
	}
}