package main

import (
	"fmt"
	"slices"
)

func main() {
	concat()
	delete()
}

func concat() {
	s1 := []string{"111"}
	s2 := []string{"222", "333", "444"}
	s3 := []string{"555", "666"}
	resp := slices.Concat(s1, s2, s3)
	fmt.Println(resp)
}

func delete() {
	s1 := []int{11, 12, 13, 14}
	s2 := slices.Delete(s1, 1, 3)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
}
