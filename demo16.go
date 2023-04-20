package main

import "fmt"

func main() {
	//这个是切片，是可以扩容的数组，是数组的升级版
	s := []int{1, 2, 3, 4}
	fmt.Println(s)
	sUpdate(s)
	fmt.Println(s)
}
func sUpdate(s []int) {
	s[0] = 100
}
