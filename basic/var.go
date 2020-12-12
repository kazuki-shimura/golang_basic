package main

import (
	"fmt"
)

func main() {

	// 変数を定義してから数値を代入
	var num1 int 
	num1 = 1
	fmt.Println(num1)

	// 変数を作成し初期値を代入
	var num2 int = 2
	fmt.Println(num2)

	// 変数を作成し初期値を代入（省略した形）
	num3 := 3
	fmt.Println(num3)

	// ハイフン、特殊文字が入っていたり数字から始まる変数は作成することができない。
	// 大文字と小文字は区別される
}