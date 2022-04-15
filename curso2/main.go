package main

import (
	"fmt"
	"golang-alura/curso2/clientes"
	"golang-alura/curso2/contas"
)

func PagarBoleto(conta verificaConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificaConta interface {
	Sacar(valor float64) string
}

func main() {

	titularBruno := clientes.Titular{
		Nome:      "Bruno",
		CPF:       "000.000.000.00",
		Profissao: "Programador"}

	contaDoBruno := contas.ContaPoupanca{
		Titular:       titularBruno,
		NumeroAgencia: 1,
		NumeroConta:   1}

	contaDoBruno.Depositar(100)
	PagarBoleto(&contaDoBruno, 60)

	fmt.Println(contaDoBruno.ObterSaldo())

}
