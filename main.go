package main

import "fmt"

type contaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *contaCorrente) sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *contaCorrente) depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "Valor inválido", c.saldo
	}
}

func (c *contaCorrente) transferir(valorDaTransferencia float64, contaDestino *contaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func main() {
	contaDaSilvia := contaCorrente{titular: "Silvia", saldo: 300}
	contaDoGustavo := contaCorrente{titular: "Gustavo", saldo: 100}

	status := contaDaSilvia.transferir(200, &contaDoGustavo)

	fmt.Println(status)
	fmt.Println(contaDoGustavo)
	fmt.Println(contaDaSilvia)

}
