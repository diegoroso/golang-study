package main

import "fmt"

// func closure() func() {
// 	x := 10
// 	var funcao = func() {
// 		fmt.Println(x)
// 	}
// 	return funcao
// }

// func main() {
// 	x := 20
// 	fmt.Println(x)

// 	imprimeX := closure()
// 	imprimeX()
// }

func closure(num int) func() int {
	var x = 10

	function := func() int {
		return x + num
	}

	return function
}

func main() {
	num := 20

	res := closure(num)
	fmt.Println("RES =>", res())
}
