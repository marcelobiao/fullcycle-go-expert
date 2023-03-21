package matematica

type Number interface {
	~int | ~float64
}

type MyNumber int

// Letra maiuscula determina se o recurso será exportado (Público)
// Letra minuscula determina se o recurso não será exportado (Privato)
// Funciona para func, structs, atributos e variáveis
func Soma[T Number](a, b T) T {
	return a + b
}

func somaNaoExportada[T Number](a, b T) T {
	return a + b
}
