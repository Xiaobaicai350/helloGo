package main

func main() {
	//可以把sum函数传进去
	a := operate(1, 2, sum)
	println(a)
	//你甚至可以在函数参数上写匿名函数
	i := operate(8, 4, func(a, b int) int {
		return a / b
	})
	println(i)
}

// 这个就是高阶函数
// 意思是可以把函数当成参数进行传递
func operate(a, b int, fun func(a, b int) int) int {
	return fun(a, b)
}
func sum(a, b int) int {
	return a + b
}
