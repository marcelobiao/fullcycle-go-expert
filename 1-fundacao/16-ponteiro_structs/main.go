package main

type Conta struct {
	Saldo int
}

func NewConta() *Conta {
	return &Conta{Saldo: 0}
}

/* func (c Conta) simular(valor int) int {
	c.Saldo += valor
	println(c.Saldo)
	return c.Saldo
} */

func (c *Conta) simular(valor int) int {
	c.Saldo += valor
	println(c.Saldo)
	return c.Saldo
}

func main() {
	conta := Conta{Saldo: 100}
	conta.simular(200)
	println(conta.Saldo)
}
