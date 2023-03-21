package main

// https://pkg.go.dev/golang.org/x/exp/constraints

// constraints
type Number interface {
	~int | ~float64
}

type MyNumber int

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// any, permite qualquer tipo, interface {}
// comparable, permite apenas itens comparáveis
func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Marcelo": 1000, "João": 2000, "Maria": 3000}
	println(Soma(m))

	m2 := map[string]float64{"Marcelo": 100.20, "João": 2000.20, "Maria": 3000.20}
	println(Soma(m2))

	// Go só aceita a maps do tipo MyNumber devido ao prefixo "~" na declaração dos tipos de dado do Number
	m3 := map[string]MyNumber{"Marcelo": 1000, "João": 2000, "Maria": 3000}
	println(Soma(m3))

	println(Compara(10, 10))
}
