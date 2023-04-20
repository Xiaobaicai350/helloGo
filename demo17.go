package main

func main() {
	println(1)
	defer println(2) //延迟执行，最先加入的最后再执行
	println(3)
	defer println(4)
	println(5)
	defer println(6)
}
