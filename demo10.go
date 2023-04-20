package main

func main() {
	var num int = 80
	switch num {
	case 90:
		println("A")
		//默认会跳出，也就是说不用加break了
		//break
	case 80:
		println("B")
		//会进行case穿透，不管下一个条件满不满足，都会执行下面的case
		fallthrough
	case 50, 60, 70:
		println("C1")
		if num == 80 {
			break //可以在这里终止穿透
		}
		println("C2")
	}
	//默认条件是true
	switch {
	case false:
		println("false")
	case true:
		println("true")
	}

}
