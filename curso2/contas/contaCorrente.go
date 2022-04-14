package contas

import (
	"golang-alura/curso2/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	if valorSaque <= c.saldo && valorSaque > 0 {
		c.saldo = c.saldo - valorSaque
		return "Saque realizado com sucesso"
	}
	return "Saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito <= 0 {
		return "Valor invalido para deposito", c.saldo
	}
	c.saldo += valorDeposito
	return "Depósito realizado com sucesso", c.saldo
}

func (c ContaCorrente) getSaldo() float64 {
	return c.saldo
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}
	return false
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
