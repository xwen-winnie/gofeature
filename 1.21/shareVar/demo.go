package main

import (
	"fmt"
)

type Item int

var prints []func()

// 共享变量例子
func main() {
	test1()
	//test2()
}

func test1() {
	var all []*Item
	items := []Item{1, 2, 3}
	for _, item := range items {
		//item := item
		all = append(all, &item)
	}
	// 打印 all 切片中的元素
	for _, item := range all {
		fmt.Println(*item)
	}
}

func test2() {
	for _, v := range []int{1, 2, 3} {
		prints = append(prints, func() { fmt.Println(v) })
	}
	for _, print := range prints {
		print()
	}
}

func test3() {
	var prints []func()
	for _, v := range []int{1, 2, 3} {
		//v := v
		prints = append(prints, func() { fmt.Println(v) })
	}
	//当这些闭包函数被调用时，它们实际上会引用循环的迭代变量v，而不是在闭包创建时的变量值, 循环在最后一次迭代后结束。
	//闭包函数捕获了v变量的引用。当闭包函数被调用时，它会通过引用找到并读取v变量的值。虽然v的生命周期已经结束，但是由于闭包函数捕获了它的引用，它仍然可以访问到v变量的值
	for _, print := range prints {
		print()
	}
}
