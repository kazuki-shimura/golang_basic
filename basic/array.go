package main

import (
	"fmt"
)


func main() {
	// 要素数を指定した配列
	a := [3]string{"sato", "suzuki", "takanashi"}
	a[0] = "tominaga"

	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])


	// 要素数を指定しない配列
	b := [...]string{"aaa", "bbb", "ccc"}

	fmt.Println(b[0])
	fmt.Println(b[1])
	fmt.Println(b[2])

	// 二次元配列
	c := [3][3]int{{1,2,3}, {4,5,6}}

	fmt.Println(c[0][0])
	fmt.Println(c[0][1])
	fmt.Println(c[2][1])
}


