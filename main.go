package main

import (
	Contas "banco/Contas"
	"fmt"
)

func main() {
	contaJose := Contas.ContaCorrente{Nome: "Jose Manoel", Saldo: 3500}
	contaMaria := Contas.ContaCorrente{Nome: "Maria", Saldo: 2500}

	status := contaJose.Transferir(800, &contaMaria)

	fmt.Println(status)
	fmt.Println(contaJose, contaMaria)

	// contaDoManoel := ContaCorrente{"Manoel da Silveira", 4596, 456548, 1500.00}
	// fmt.Println(contaDoManoel.sacar(650.55), contaDoManoel)
	// fmt.Println("Conta do manoel: ", contaDoManoel)

	// var contaJose ContaCorrente
	// contaJose.nome = "jose manoel"
	// contaJose.saldo = 2500
	// // saque := 200.00

	// valorDeposito := 300
	// fmt.Println(contaJose.sacar(500), contaJose)
	// depositoMsg, saldo := contaJose.depositar(float64(valorDeposito))
	// fmt.Println(depositoMsg, "No valor de: ", valorDeposito, "O saldo é: ", saldo)

}
