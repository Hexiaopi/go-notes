package stack

import "errors"

type Stack []interface{}

func(stack *Stack)Push(i interface{}){
	*stack=append(*stack,i)
}

func(stack *Stack)Pop()(interface{},error){
	if len(*stack)==0{
		return nil,errors.New("Stack already empty")
	}
	i:=(*stack)[len(*stack)-1]
	*stack=(*stack)[:len(*stack)-1]
	return i,nil
}

func(stack *Stack)Top()(interface{},error){
	if len(*stack)==0{
		return nil,errors.New("Stack already empty")
	}
	return (*stack)[len(*stack)-1],nil
}

func(stack *Stack)Len()int{
	return len(*stack)
}

func (stack *Stack) Cap() int {
	return cap(*stack)
}

func (stack *Stack)IsEmpty() bool{
	return len(*stack)==0
}