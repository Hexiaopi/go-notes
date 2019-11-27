package duck_1

import "fmt"


type FlyBehavior interface {
	Fly()
}

type FlyWithWings struct {

}

func (f FlyWithWings)Fly(){
	fmt.Println("fly with wings")
}

type FlyNoWay struct {

}

func (f FlyNoWay)Fly(){
	//do nothing
}

type QuackBehavior interface {
	Quack()
}

type Quack struct {

}

func (q Quack)Quack(){
	fmt.Println("gua gua...")
}

type Squeak struct {

}

func (s Squeak)Quack(){
	fmt.Println("zi zi...")
}

type MuteQuack struct {

}

func (m MuteQuack)Quack(){
	//do nothing
}

type Duck struct {
	QuackBehavior
	FlyBehavior
}

func(d Duck)performQuack(){
	d.Quack()
}

func (d Duck)performFly(){
	d.Fly()
}

func(d Duck)Swim(){
}

func(d Duck)Display(){
	fmt.Println("Duck")
}

type MallardDuck struct {
	Duck
}

func NewMallardDuck()MallardDuck{
	return MallardDuck{
		Duck{
			QuackBehavior:Quack{},
			FlyBehavior:FlyWithWings{},
		},
	}
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


func (rd RubberDuck)Display(){
	fmt.Println("Rubber Duck")
}