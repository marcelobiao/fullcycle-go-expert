package main

import (
	"fmt"

	myMath "github.com/marcelobiao/fullcycle-go-expert/5-packaging/1/math"
)

func main() {
	m := myMath.Math{A: 5, B: 5}
	fmt.Println(m.Add())
}
