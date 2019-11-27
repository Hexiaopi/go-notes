package duck_0


func ExampleDuck_Quack() {
	duck:=Duck{}
	duck.Quack()
	//Output:
	//gua gua...
}

func ExampleDuck_Display() {
	duck:=Duck{}
	duck.Display()
	//Output:
	//Duck
}

func ExampleMallardDuck_Display() {
	duck:=MallardDuck{}
	duck.Display()
	//Output:
	//Mallard Duck
}

func ExampleRedheadDuck_Display() {
	duck:=RedheadDuck{}
	duck.Display()
	//Output:
	//Redhead Duck
}

func ExampleRubberDuck_Display() {
	duck:=RubberDuck{}
	duck.Display()
	//Output:
	//Rubber Duck
}

func ExampleRubberDuck_Quack() {
	duck:=RubberDuck{}
	duck.Quack()
	//Output:
	//zi zi...
}