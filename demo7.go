package main

import "fmt"

func main() {
	v1 := 'a'
	v2 := "a"
	fmt.Printf("%T,%d", v1, v1)
	println()
	fmt.Printf("%T,%s", v2, v2)

}
