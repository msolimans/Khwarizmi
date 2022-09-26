package evaluatereversepolishnotation

import (
	"container/list"
	"strconv"
)

func evalRPN(tokens []string) int {
    stack := list.New() 

	for _, t := range tokens {
		if !isOperand(t) {//push to stack 
			stack.PushFront(t)
		}else {//otherwise pop 2 and calc then save result back to stack 
			
			//get 2 operatos for this operand to calc 
			val1 := stack.Front() 
			stack.Remove(val1)
			val2 := stack.Front() 
			stack.Remove(val2)
			
			//pushback result 
			res := calc(val2.Value.(string), val1.Value.(string), t)
			stack.PushFront(res) 
		}
	}

	resAsStr := stack.Back().Value.(string)
	res,_ := strconv.Atoi(resAsStr)
	return res 
}

var mp = map[string]bool{"+": true, "-":true, "*":true, "/":true }

func calc(operator1, operator2, operand string) string {
	op1,_:= strconv.Atoi(operator1)
	op2,_ := strconv.Atoi(operator2)
	
	switch operand {
	case "+":
		return strconv.Itoa(op1+op2)
	case "-":
		return strconv.Itoa(op1-op2)
	case "*":
		return strconv.Itoa(op1*op2)
	case "/":
		return strconv.Itoa(op1/op2)
	default:
		return ""
	}
}
func isOperand(s string) bool {
	_,exists := mp[s]

	return exists 
}