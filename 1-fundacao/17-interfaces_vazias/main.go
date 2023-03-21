package main

import "fmt"

type x interface{}

func main() {
	var x interface{} = 10
	var y interface{} = "hello, World!"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("Tipo: %T - valor %v\n", t, t)
}
