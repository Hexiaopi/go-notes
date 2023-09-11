package flagdemo

import (
	"flag"
	"fmt"
)

func ExampleFlagStringPointPrint() {
	flag.CommandLine.Parse([]string{"-str-a", "stra"})
	fmt.Println(*flagStrA)
	// Output:
	// stra
}

func ExampleFlagStringPrint() {
	flag.CommandLine.Parse([]string{"-str-b", "strb"})
	fmt.Println(flagStrB)
	// Output:
	// strb
}

func ExampleFlagIntPointPrint() {
	flag.CommandLine.Parse([]string{"-int", "10"})
	fmt.Println(*flagInt)
	// Output:
	// 10
}

func ExampleBoolPointPrint() {
	flag.CommandLine.Parse([]string{"-bool=True"})
	fmt.Println(*flagBool)
	// Output:
	// true
}

func ExampleBoolPointNoticPrint() {
	flag.CommandLine.Parse([]string{"-bool", "false"})
	fmt.Println(*flagBool)
	// Output:
	// true
}

func ExampleBoolPointEmptyPrint() {
	flag.CommandLine.Parse([]string{"-bool"})
	fmt.Println(*flagBool)
	// Output:
	// true
}

func ExampleFlagDurationPointPrint() {
	flag.CommandLine.Parse([]string{"-duration", "1m1s"})
	fmt.Println(*flagDuration)
	// Output:
	// 1m1s
}

func ExampleStructPrint() {
	flag.CommandLine.Parse([]string{"-url", "https://golang.org/pkg/flag/"})
	fmt.Println(u)
	// Output:
	// https://golang.org/pkg/flag/
}
