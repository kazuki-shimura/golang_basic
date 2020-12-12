package main

import (
	"fmt"
	"reflect"
)

func main(){

	// 数値型
	num1 := 10
	var num2 int = 1234567890
	num3 := 1.23
	var num4 float64 = 1.234567

	fmt.Print(reflect.TypeOf(num1))
	fmt.Println(num1)
	fmt.Print(reflect.TypeOf(num2))
	fmt.Println(num2)
	fmt.Print(reflect.TypeOf(num3))
	fmt.Println(num3)
	fmt.Print(reflect.TypeOf(num4))
	fmt.Println(num4)

	// 文字列型
	var string_a string = "行ってきます！"
	string_b := "ただいま！"
	
	fmt.Print(reflect.TypeOf(string_a))
	fmt.Println(string_a)
	fmt.Print(reflect.TypeOf(string_b))
	fmt.Println(string_b)

	// ブール型
	var bool_data bool = true
	a := 1
	b := 5
	var num_bool bool = a > b

	fmt.Print(reflect.TypeOf(bool_data))
	fmt.Println(bool_data)
	fmt.Print(reflect.TypeOf(num_bool))
	fmt.Println(num_bool)
}

