package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// 存储
	m.Store("1", "one")
	m.Store("2", "two")

	// 查询
	one, ok := m.Load("1")
	if !ok {
		fmt.Println("key 1 not found")
	} else {
		fmt.Println(one.(string))
	}

	// 遍历
	m.Range(func(key, value any) bool {
		fmt.Println("key:", key, "value:", value)
		return true
	})

	// 删除
	m.Delete("1")
	one, ok = m.Load("1")
	if !ok {
		fmt.Println("key 1 not found")
	} else {
		fmt.Println(one.(string))
	}

	// 查询，不存在的情况下存储默认值
	two, loaded := m.LoadOrStore("2", "two")
	fmt.Println("three:", two, "loaded:", loaded)
	three, loaded := m.LoadOrStore("3", "three")
	fmt.Println("three:", three, "loaded:", loaded)

	// 查询并删除
	three, loaded = m.LoadAndDelete("3")
	fmt.Println("three:", three, "loaded:", loaded)
	four, loaded := m.LoadAndDelete("4")
	fmt.Println("four:", four, "loaded:", loaded)

	// 交换
	two, loaded = m.Swap("2", "double")
	fmt.Println("two:", two, "loaded:", loaded)
	two, ok = m.Load("2")
	fmt.Println(two, ok)

	// 交换，仅当旧值相同时才进行交换
	result := m.CompareAndSwap("2", "two", "double")
	fmt.Println(result)
	result = m.CompareAndSwap("2", "double", "two")
	fmt.Println(result)
}
