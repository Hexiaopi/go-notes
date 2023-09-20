package pflagdemo

import (
	"fmt"
	"strings"

	flag "github.com/spf13/pflag"
)

var name = flag.StringP("name", "n", "zhangsan", "Input Your Name")
var age = flag.IntP("age", "a", 30, "Input Your Age")
var gender = flag.StringP("gender", "g", "male", "Input Your Gender")
var married = flag.BoolP("married", "m", false, "Input Are You Merryied")
var desc = flag.StringP("desc", "d", "", "Input Description")
var badFlag = flag.StringP("badflag", "b", "just for test", "Input badflag")

func wordSepNormalizeFunc(f *flag.FlagSet, name string) flag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return flag.NormalizedName(name)
}

func main() {
	// 设置标准化参数名称的函数
	flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)

	// 为 age 参数设置 NoOptDefVal
	flag.Lookup("age").NoOptDefVal = "25"

	// 把 badflag 参数标记为即将废弃的，请用户使用 desc 参数
	flag.CommandLine.MarkDeprecated("badflag", "please use --desc instead")

	// 把 badflag 参数的 shorthand 标记为即将废弃的，请用户使用 desc 的 shorthand 参数
	flag.CommandLine.MarkShorthandDeprecated("badflag", "please use -d instead")

	// 在帮助文档中隐藏参数 gender
	flag.CommandLine.MarkHidden("badflag")

	// 把用户传递的命令行参数解析为对应变量的值
	flag.Parse()

	fmt.Println("name:", *name)
	fmt.Println("age:", *age)
	fmt.Println("gender:", *gender)
	fmt.Println("married:", *married)
	fmt.Println("desc:", *desc)
}
