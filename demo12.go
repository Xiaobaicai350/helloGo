package main

func main() {
	str := "hello"
	println(len(str)) //5
	println(str[0])   //104
	println(str[1])   //101
	//i是下标 v是值
	for i, v := range str {
		println(i, v)
	}

}
