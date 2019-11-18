package main

import "fmt"

type Component interface {
	Calc() int
}

type ConcreteComponent struct{}

func (*ConcreteComponent) Calc() int {
	return 0
}

// MulDecorator 装饰者
type MulDecorator struct {
	Component
	num int
}

func WarpMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

func (d *MulDecorator) Calc() int {
	return d.Component.Calc() * d.num
}

// AddDecorator 装饰者
type AddDecorator struct {
	Component
	num int
}

func WarpAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

func (d *AddDecorator) Calc() int {
	return d.Component.Calc() + d.num
}

func main() {
	var c Component = &ConcreteComponent{}	// 0
	c = WarpAddDecorator(c, 10)	// 0+10 具有了和的行为
	c = WarpMulDecorator(c, 8)		// 10*8	具有了积的行为
	res := c.Calc()

	fmt.Printf("res %d\n", res)
}
