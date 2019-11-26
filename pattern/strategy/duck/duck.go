package duck

import "fmt"

type Duck struct {}

func(d Duck)Quack(){
	fmt.Println("gua gua...")
}

func(d Duck)Swim(){
}

func(d Duck)Display(){
	fmt.Println("Duck")
}

type MallardDuck struct {
	Duck
}

func (md MallardDuck)Display(){
	fmt.Println("Mallard Duck")
}

type RedheadDuck struct {
	Duck
}

func (rd RedheadDuck)Display(){
	fmt.Println("Redhead Duck")
}

type RubberDuck struct{
	Duck
}

func (rd RubberDuck)Quack(){
	fmt.Println("zi zi...")
}

func (rd RubberDuck)Display(){
	fmt.Println("Rubber Duck")
}
