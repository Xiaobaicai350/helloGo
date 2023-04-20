package main

//一个外层函数中，有内层函数，该内层函数中，会操作外层函数的局部变量且该外层函数的返回值就是这个内层函数。
//这个内层函数和外层函数的局部变量，统称为闭包结构

// 局部变量的生命周期就会发生改变，正常的局部变量会随着函数的调用而创建，
// 随着函数的结束而销毁但是闭包结构中的外层函数的局部变量并不会随着外层函数的结束而销毁，因为内层函数还在继续使用
func main() {
	//这里获得内部的fun函数
	method1 := increment()
	println(method1()) //1
	println(method1()) //2
	method2 := method1
	println(method2()) //3
	println(method2()) //4
	method3 := increment()
	println(method3()) //1
	println(method3()) //2

}
func increment() func() int {
	//局部变量i
	i := 0
	//定义一个函数，让i自增并返回
	fun := func() int {
		i++
		return i
	}
	return fun
}
