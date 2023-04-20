package main

func main() {
	f6(1, 2, 3, 4)
}

// - 无参无返回值函数
func f1() {
	println("1")
}

// - 有一个参数的函数
func f2(a int) {
	println(a)
}

// - 有多个参数的函数
func f3(a, b int) {
	println(a, b)
}

// - 有一个返回值的函数
func f4() int {
	return 1
}

// - 有多个返回值的函数
// f5 函数注释（可以理解为Java中的文档注释）
func f5() (int, int) {
	return 1, 2
}

// f6 可变参数
func f6(nums ...int) {
	i := len(nums)
	println(i)
}
