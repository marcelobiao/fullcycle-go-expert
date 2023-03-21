package main

import "fmt"

type ID int

func main() {
	var f ID = 1
	fmt.Printf("O tipo é %T", f)
}
