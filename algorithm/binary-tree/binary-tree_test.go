package binary_tree

var (
	four = Node{
		Left:  nil,
		Data:  4,
		Right: nil,
	}
	five = Node{
		Left:  nil,
		Data:  5,
		Right: nil,
	}
	two = Node{
		Left:  &four,
		Data:  2,
		Right: &five,
	}
	sixth = Node{
		Left:  nil,
		Data:  6,
		Right: nil,
	}
	seven = Node{
		Left:  nil,
		Data:  7,
		Right: nil,
	}
	three = Node{
		Left:  &sixth,
		Data:  3,
		Right: &seven,
	}
	one = Node{
		Left:  &two,
		Data:  1,
		Right: &three,
	}
)

func ExamplePreOrder() {
	PreOrder(&one)
	// Output:
	// 1 2 4 5 3 6 7
}

func ExampleInOrder() {
	InOrder(&one)
	// Output:
	// 4 2 5 1 6 3 7
}

func ExampleAftOrder() {
	AftOrder(&one)
	// Output:
	// 4 5 2 6 7 3 1
}

func ExampleLayerOrder() {
	LayerOrder(&one)
	// Output:
	// 1 2 3 4 5 6 7
}
