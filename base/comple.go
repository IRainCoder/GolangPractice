/*
复合数据结构
数组
切片
映射
结构体
*/
package main

import (
	"fmt"
)


func main() {
	// 数组
	var arr = [...]int{1, 2, 3}
	// fmt.Println(arr)
	for i, v := range arr {
		fmt.Println(i, v)
	}

	// 	切片
	var slice = []int{1, 2, 3}
	slice = append(slice, 4)
	for i, v := range slice {
		fmt.Println(i, v)
	}

	// 映射
	var ages = map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	for name, age := range ages {
		fmt.Println(name, age)
	}
	fmt.Println(ages["Alice1"])

	// 结构体
	type Person struct {
		Name string
		Age  int
	}
	var p = Person{
		Name: "Alice",
		Age:  25,
	}
	fmt.Println(p.Name, p.Age)

}