package main

import "fmt"

func main() {
	// clear 演示
	m := map[string]string{"你好": "世界", "世界": "你好"}
	fmt.Printf("m1: %v, len: %d\n", m, len(m))
	clear(m)
	fmt.Printf("m2: %v, len: %d\n", m, len(m))

	// min max demo
	//var x, y int
	//x = 3
	//y = 4
	////m := min(x)             // m == x
	//m := min(x, y)          // m 是 x 和 y 中较小的那个
	//m := max(x, y, 10) // m 是 x 和 y 中较大的一个，但至少是10
	c := max(1.0, 20.0, 10) // c == 10.0（浮点类型）
	//f := max(0, float32(x), float32(y)) // f 的类型是 float32
	//var s []string
	//_ = min(s...)              // 无效：不允许使用 slice 参数
	t := max("vsdfsa", "zoo", "yar") // t == "foo" (string 类型)
	fmt.Println("m is:", t)
	fmt.Println("c is:", c)

}
