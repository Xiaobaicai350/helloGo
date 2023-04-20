package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	//这里就和Java不一样了，这里是进行了值拷贝
	//也就是说把实参的值赋给了形参，并没有改变原来的值
	update(arr)
	fmt.Println(arr)
}
func update(arr [4]int) {
	arr[0] = 100
}
