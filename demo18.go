package main

func main() {
	var num int = 10
	defer f(num) //可以发现，程序执行到这一步已经把num的值传进去了（值传递）
	num++
	println(num) //先执行这个
}
func f(num int) {
	println(num)
}
