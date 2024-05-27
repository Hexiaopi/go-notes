package main

import (
	"bytes"
	"fmt"
	"strings"
)

func PrintString() {
	s := "hello,world"
	fmt.Println(len(s))
}

func PrintStringCompare() {
	s1 := "世界和平"
	s2 := "世界" + "和平"
	fmt.Println(s1 == s2) //true

	s1 = "Go"
	s2 = "c"
	fmt.Println(s1 != s2) //true

	s1 = "12345"
	s2 = "23456"
	fmt.Println(s1 < s2)  //true
	fmt.Println(s1 <= s2) //true

	s1 = "12345"
	s2 = "123"
	fmt.Println(s1 > s2)  //true
	fmt.Println(s1 >= s2) //true
}

func PrintStringByte(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Println(i, string(str[i]))
	}
}

func PrintStringRune(str string) {
	for i, c := range str {
		fmt.Println(i, string(c))
	}
}

var sl []string = []string{
	"Rob Pike ",
	"Robert Griesemer ",
	"Ken Thompson ",
}

func concatStringByOperator(sl []string) string {
	var s string
	for _, v := range sl {
		s += v
	}
	return s
}

func concatStringBySprintf(sl []string) string {
	var s string
	for _, v := range sl {
		s = fmt.Sprintf("%s%s", s, v)
	}
	return s
}

func concatStringByJoin(sl []string) string {
	return strings.Join(sl, "")
}

func concatStringByStringsBuilder(sl []string) string {
	var b strings.Builder
	for _, v := range sl {
		b.WriteString(v)
	}
	return b.String()
}

func concatStringByStringsBuilderWithInitSize(sl []string) string {
	var b strings.Builder
	b.Grow(64)
	for _, v := range sl {
		b.WriteString(v)
	}
	return b.String()
}

func concatStringByBytesBuffer(sl []string) string {
	var b bytes.Buffer
	for _, v := range sl {
		b.WriteString(v)
	}
	return b.String()
}

func concatStringByBytesBufferWithInitSize(sl []string) string {
	buf := make([]byte, 0, 64)
	b := bytes.NewBuffer(buf)
	for _, v := range sl {
		b.WriteString(v)
	}
	return b.String()
}
