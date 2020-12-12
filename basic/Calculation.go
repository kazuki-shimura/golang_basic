package main

import (
	"fmt"
)

func main(){
	// 四則演算
	x := 10
	y := 2

	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	// 関係演算子
	fmt.Println(x < y)

	// 等価
	fmt.Println(x == y)
	fmt.Println(x != y)

	// 論理演算子
	fmt.Println(x > 5 && x <= 8)
	fmt.Println(x > 5 || x <= 8)

	// 複合代入演算子
	x += 20
	y += x
	fmt.Println(x)
	fmt.Println(y)

	// インクリメントデクリメント
	x++
	y--
	fmt.Println(x)
	fmt.Println(y)
}