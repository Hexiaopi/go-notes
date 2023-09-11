package flagdemo

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"time"
)

// Example1：flag.String 返回指针类型字符串
var flagStrA = flag.String("str-a", "str-a", "flag test *string")

func FlagStringPointPrint() {
	fmt.Println(*flagStrA)
}

// Example2：如果期望定义字符串类型，可以通过init函数初始化
var flagStrB string

func init() {
	// flag.StringVar可以绑定已有的变量
	flag.StringVar(&flagStrB, "str-b", "str-b", "flag test string")
}

// Example3：int类型变量
var flagInt = flag.Int("int", 1, "flag test *int")

// Example4: bool类型变量，注意使用
var flagBool = flag.Bool("bool", false, "flag test *bool")

// Example5：duration类型变量
var flagDuration = flag.Duration("duration", time.Minute*1, "flag test *duration")

// Example6: 自定义类型
type URLValue struct {
	URL *url.URL
}

func (v URLValue) String() string {
	if v.URL != nil {
		return v.URL.String()
	}
	return ""
}

func (v *URLValue) Set(s string) error {
	if u, err := url.Parse(s); err != nil {
		return err
	} else {
		v.URL = u
	}
	return nil
}

var u URLValue

func init() {
	flag.Var(&u, "url", "flag test self struct")
}
