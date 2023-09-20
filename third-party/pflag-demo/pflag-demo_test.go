package pflagdemo

import (
	"fmt"
	"testing"

	"github.com/spf13/pflag"
)

func ExampleFlagLongStringPoint() {
	pflag.CommandLine.Parse([]string{"--str-a", "stra"})
	fmt.Println(*flagStrA)
	// Output:
	// stra
}

func ExampleFlagShortStringPoint() {
	pflag.CommandLine.Parse([]string{"-b", "strb"})
	fmt.Println(*flagStrB)
	// OutPut:
	// strb
}

func ExampleFlagString() {
	pflag.CommandLine.Parse([]string{"--str-c", "strc", "-d", "strd"})
	fmt.Println(flagStrC, flagStrD)
	// Output:
	// strc strd
}

func TestFlagDeprecated(t *testing.T) {
	pflag.CommandLine.Parse([]string{"--badflag", "deprecated", "-b", "bbb"})
	t.Log(pflag.CommandLine.FlagUsages())
}

func TestFlagDefault(t *testing.T) {
	tests := []struct {
		name      string
		arguments []string
		want      int
	}{
		{name: "没传入参数和值", arguments: []string{""}, want: 30},
		{name: "传入参数和值", arguments: []string{"--age=10"}, want: 10},
		{name: "传入参数,没传值", arguments: []string{"--age"}, want: 25},
	}
	for _, test := range tests {
		pflag.CommandLine.Parse(test.arguments)
		if *age != test.want {
			t.Fatalf("%s pflag parse want:%d,but got:%d", test.name, test.want, *age)
		}
	}
}
