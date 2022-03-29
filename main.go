package main

import (
	"fmt"

	// Clientes "github.com/domjesus/banco/clientes"
	Contas "github.com/domjesus/banco/contas"
	// Clientes "github.com/domjesus/banco/clientes"
)

func main() {
	contaExemplo := Contas.ContaPoupanca{}

	contaExemplo.Depositar(1100)
	pagarBoleto(&contaExemplo, 700)

	fmt.Println(contaExemplo.GetSaldo())
	// clienteJose := Clientes.Titular{"Jose Manoel da Silva", "123.123.123-55", "Dev master"}
	// contaJose := Contas.ContaCorrente{clienteJose, 1234, 321654, 100.00}
	// contaJose := Contas.ContaCorrente{Titular: Clientes.Titular{Nome: "Nome do titular", CPF: "987654987", Profissao: "Dev"}, Saldo: 1500., NumeroAgencia: 1500, NumeroConta: 1234}

	// fmt.Println(contaJose)
	// cliente := Clientes.Titular{Nome: "NOSE MANOEL CLIENTE", CPF: "45654987"}
	// fmt.Println(cliente)

	// contaJose := Contas.ContaCorrente{Nome: "Jose Manoel", Saldo: 3500}
	// contaMaria := Contas.ContaCorrente{Nome: "Maria", Saldo: 2500}

	// status := contaJose.Transferir(800, &contaMaria)

	// fmt.Println(status)
	// fmt.Println(contaJose, contaMaria)

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

type verificarConta interface {
	Sacar(valor float64) string
}

func pagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}
