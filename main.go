package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func pagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	pagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())
	fmt.Println(clientes.Titular{"Denis", "123.3459.394", "Desenvolvedor"})

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(500)
	pagarBoleto(&contaDaLuisa, 400)
	fmt.Println(contaDaLuisa.ObterSaldo())
}
