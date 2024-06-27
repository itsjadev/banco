package main

import "fmt"

type contaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDaJade := contaCorrente{"Jade", 589, 123456, 125.5}
	contaDaBruna := contaCorrente{"Bruna", 222, 11112, 200}
	fmt.Println(contaDaJade)
	fmt.Println(contaDaBruna)
}
