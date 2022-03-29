package Contas

import Clientes "github.com/domjesus/banco/clientes"

type ContaPoupanca struct {
	Titular  Clientes.Titular
	operacao string
	saldo    float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) string {
	podeSacar := valorSaque <= c.saldo && valorSaque > 0

	if podeSacar {
		c.saldo -= valorSaque
		return "Saque de realizado com sucesso!"
	} else {
		return ("saldo insuficiente ou número incorreto!")
	}

	// return " alguma coisa"

	// fmt.Println(podeSacar)

}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito <= 0 {
		return "Para depositar deve ser informado número positivo!", c.saldo
	}

	c.saldo += valorDeposito
	return "Depósito feito com sucesso!", c.saldo
}

func (c *ContaPoupanca) GetSaldo() float64 {
	return c.saldo
}
