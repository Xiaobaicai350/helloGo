package main

func test() (int, int) {
	return 100, 200
}

func main() {
	//a, b := test()
	//println(a)
	//println(b)
	//_叫空白标识符，给这个符号赋值的数据会被抛弃
	//这个也叫匿名变量
	//匿名变量不占用内存空间。不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。
	a, _ := test()
	println(a)
}
