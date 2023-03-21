package main

func main() {
	// Memória -> Endereço -> Valor

	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a
	println(*b)

	a1 := 10
	b1 := 20
	println(soma(a1, b1))
	println(a1)
	println(somaAlterada(&a1, &b1))
	println(a1)
}

func soma(a, b int) int {
	a = 50
	return a + b
}

func somaAlterada(a, b *int) int {
	*a = 50
	return *a + *b
}
