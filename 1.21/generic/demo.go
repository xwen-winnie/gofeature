package main

import "fmt"

// MapKeys 泛型函数，用于获取一个map 的所有键, 返回一个键的切片
func MapKeys[K comparable, V any](m map[K]V) []K {
	// 创建一个切片r初始容量为map 的长度, 用于存放所有键
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	// 返回包含所有键的切片
	return r
}

// any 关键字是一个预定义的类型约束，声明后将允许任何类型用作类型实参，并且允许函数使用用于任何类型的操作
// 泛型结构体
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("keys:", MapKeys(m))

	// 显示指定类型参数，再次调用MapKeys 函数，与上面等效
	//_ = MapKeys[int, string](m)
	fmt.Println("keys:", MapKeys[int, string](m))

	// 创建空的整数类型的链表
	lst := List[int]{}

	// 向链表添加元素
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	// 调用GetAll 方法,获取链表中的所有元素 打印
	fmt.Println("list:", lst.GetAll())
}
