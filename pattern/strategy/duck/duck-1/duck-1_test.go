package duck_1

import "testing"

func TestMallardDuck_Fly(t *testing.T) {
	md:=NewMallardDuck()
	md.Fly()
	//fly with wings
}

func TestRedheadDuck_Quack(t *testing.T) {
	md:=NewMallardDuck()
	md.Quack()
	//gua gua...
}
