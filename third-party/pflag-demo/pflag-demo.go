package pflagdemo

import (
	"github.com/spf13/pflag"
)

var flagStrA = pflag.String("str-a", "str-a", "long flag test *string")

var flagStrB = pflag.StringP("str-b", "b", "str-b", "long and short flag test *string")

var flagStrC string
var flagStrD string

func init() {
	pflag.StringVar(&flagStrC, "str-c", "str-c", "long flag test string")
	pflag.StringVarP(&flagStrD, "str-d", "d", "str-d", "long and short flag test string")
}

var badFlag = pflag.StringP("badflag", "f", "just for test", "Input badflag")

func init() {
	pflag.CommandLine.MarkDeprecated("badflag", "please use --desc instead")
	pflag.CommandLine.MarkShorthandDeprecated("str-b", "please use --str-b instead")
}

var age = pflag.IntP("age", "a", 30, "Input Your Age")

func init() {
	pflag.Lookup("age").NoOptDefVal = "25"
}
