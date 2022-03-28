package Contas

import "math"

type ContaCorrente struct {
	Nome          string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) sacar(valorSaque float64) string {
	podeSacar := valorSaque <= c.Saldo && valorSaque > 0

	if podeSacar {
		c.Saldo -= math.Abs(valorSaque)
		return "Saque de realizado com sucesso!"
	} else {
		return ("Saldo insuficiente ou número incorreto!")
	}

	// return " alguma coisa"

	// fmt.Println(podeSacar)

}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito <= 0 {
		return "Para depositar deve ser informado número positivo!", c.Saldo
	}

	c.Saldo += valorDeposito
	return "Depósito feito com sucesso!", c.Saldo
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {

	if valorTransferencia <= c.Saldo && valorTransferencia > 0 {
		c.Saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
		//FAZ TRANSFERENCIA
	} else {
		return false
	}

}
