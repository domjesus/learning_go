package Contas

import (
	"math"

	"github.com/domjesus/banco/clientes"
)

type ContaCorrente struct {
	Titular                    Clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque <= c.saldo && valorSaque > 0

	if podeSacar {
		c.saldo -= math.Abs(valorSaque)
		return "Saque de realizado com sucesso!"
	} else {
		return ("saldo insuficiente ou número incorreto!")
	}

	// return " alguma coisa"

	// fmt.Println(podeSacar)

}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito <= 0 {
		return "Para depositar deve ser informado número positivo!", c.saldo
	}

	c.saldo += valorDeposito
	return "Depósito feito com sucesso!", c.saldo
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {

	if valorTransferencia <= c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
		//FAZ TRANSFERENCIA
	} else {
		return false
	}
}

func (c *ContaCorrente) GetSaldo() float64 {
	return c.saldo
}
