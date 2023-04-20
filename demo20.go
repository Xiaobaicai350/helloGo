package main

func main() {
	//第一种调用方式
	func() {
		println("我是匿名函数")
	}()
	//第二种调用方式
	funcDemo := func() {
		println("我也是匿名函数")
	}
	funcDemo()
	//带参数的匿名函数
	func(a, b int) {
		println("我是带参数的匿名函数")
		println(a, b)
	}(1, 2)
	//带返回值的匿名函数
	sum := func(a, b int) int {
		return a + b
	}(1, 2)
	println(sum)
}
