package main

import "fmt"

func main() {
	//打印出来的函数类型
	//func(int, int)
	fmt.Printf("%T\n", func1)
	//定义一个函数类型的变量
	var func2 func(int, int)
	//给函数类型的变量赋值
	func2 = func1
	//调用这个函数
	func2(1, 2)
}
func func1(a, b int) {
	fmt.Println(a, b)
}
