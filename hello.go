// 在Go语言里，命名为main的包具有特殊的含义。
// Go语言的编译程序会试图把这种名字的包编译为二进制可执行文件。
// 所有用Go语言编译的可执行程序都必须有一个名叫 main的包。一个可执行程序有且仅有一个 main包
package main

import "fmt"

func main() {
	//声明变量用var
	//name 是变量名
	//string/int 代表变量的类型
	var name string = "胥天昊"
	var age int = 18
	var (
		//如果没有显式赋值，默认有初始值
		//string 		空字符串
		//int 			0
		//bool 			false
		//切片/函数/指针   nil
		sex  string = "男"
		addr string = "河南"
	)
	//甚至可以这样定义
	var a, b, c int = 1, 2, 3
	//还可以这样，编译器默认推导变量是什么类型
	//这个名字叫短变量声明并初始化
	n1 := 10
	n2 := "变量"

	println("hello " + name)
	println(age)
	println(sex, addr)
	println(a, b, c)
	//%T是个占位符，可以打印出来变量的类型
	fmt.Printf("%T,%T", n1, n2)
}
