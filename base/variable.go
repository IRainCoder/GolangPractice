/*
变量

1. 变量名称
- 字母、数字、下划线组成，数字不能开头，并区分大小写。
- 不能使用Go语言的关键字，即以下25个关键字：
break		default      func         interface    select
case		defer        go           map          struct
chan		else         goto         package      switch
const		fallthrough  if           range        type
continue	for          import       return       var
- 不要使用内置的预声明的常量、类型和函数，如以下内容：
常量：true false iota nil
类型：int int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 uintptr
	float32 float64 complex128 complex64
	bool byte rune string error
函数：make len cap new append copy close delete
	complex real imag panic recover

2. 变量声明
- 通用形式：var 变量名称 类型 = 表达式
- 初始值：Go里面不存在未初始化的变量，声明变量时如果不指定初始值，变量会被自动赋值为其类型的零值。包级别的变量初始化在main开始之前进行。
以下是一些零值说明：
对于数字类型，零值是0；
对于布尔类型，零值是false；
对于字符串类型，零值是空字符串""；
对于接口和引用类型（包括slice、指针、map、chan和函数），零值是nil。
对于数组或结构体等聚合类型，零值是每个元素或字段都是对应该类型的零值。

3. 短变量声明
- 通用形式：变量名称 := 表达式
- 特性1: 短变量声明语句中必须至少要声明一个新的变量，否则会导致编译错误。
- 特性2: 短变量不能作为包级别的变量使用。

4. new函数创建变量
- 通用形式：变量名称 := new(类型)
- 函数返回值：new函数创建一个未命名的变量，并返回具有唯一(例外：两个变量的类型不携带任何信息且是零值，比如strruct{})地址的变量。

5. 指针变量
- 通用形式1：var 变量名称 *类型 = &变量
- 通用形式2：变量名称 := &变量
- 作用：在无需知道变量的具体类型的情况下，间接地访问或更新变量的值。
- 特性1: 变量是存储值的地方，指针是存储变量地址的地方。所有的变量都有地址，但不是所有的值都有地址。
- 特性2: 两个指针当且仅当它们指向同一个变量或者两者都是nil时才相等。

6. 变量的生命周期
- 定义：变量的生命周期是指在程序运行期间变量存在的时间段。
- 特性1: 包级别的变量的生命周期是整个程序的执行时间。

7. 变量的作用域
- 定义：变量的作用域是指程序源代码中定义这个变量的作用区域。
- 规则：
如果一个变量是在函数内部声明的，那么它只能在该函数内部访问。
如果一个变量是在函数外部声明的，那么它可以在整个包内访问。
如果一个变量是在函数外部声明且以大写字母开头，那么它可以在整个程序内访问。

8. 赋值
- 通用形式1：变量名称 = 表达式
- 通用形式2：变量名称1, 变量名称2, ... = 表达式1, 表达式2, ...
- 通用形式3：变量名称 关系运算符= 表达式  (v += 1)
- 特性1: 赋值只有在值对变量类型是可赋值时才合法。

*/
package main

import "fmt"


// 包级别变量
var pi float64


func main() {
	pi = 3.14
	r := 5.0

	// 面积
	m := new(float64)

	// 短变量, 计算圆周
	c := pi * r * 2.0
	fmt.Println(c)

	// 指针，计算面积
	r2 := r * r
	*m = pi * r2
	fmt.Println(*m)
	fmt.Printf("变量地址：%v", m)
}
