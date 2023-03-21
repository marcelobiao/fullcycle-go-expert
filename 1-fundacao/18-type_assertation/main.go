package main

type x interface{}

func main() {
	var minhaVar interface{} = "Marcelo"
	println(minhaVar.(string))

	res, ok := minhaVar.(string)
	if !ok {
		println("error")
	}
	println(res)

	res2, ok := minhaVar.(int)
	if !ok {
		println("error")
	}
	println(res2)
}
