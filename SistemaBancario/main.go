package main

import (
	"fmt"
)

type Operacao struct {
	Tipo  string
	Valor float64
}

var operacoes []Operacao
var saldo float64 = 0
var numero_saques int = 0

func depositar() {
	var valorDep float64
	fmt.Println("DEPÓSITO")
	fmt.Print("Informe o valor a ser depositado: ")
	fmt.Scan(&valorDep)
	if valorDep <= 0 {
		fmt.Println("Não permitido depósito de valor negativo")
	} else {
		operacoes = append(operacoes, Operacao{Tipo: "Depósito", Valor: valorDep})
		saldo += valorDep
		fmt.Printf("Deposito realizado no valor de R$ %.2f\n", valorDep)
		fmt.Printf("Seu saldo é R$ %.2f\n", saldo)
	}
}

func sacar() {
	var valorSaq float64
	fmt.Println("SAQUE")
	fmt.Print("Informe o valor a ser sacado: ")
	fmt.Scan(&valorSaq)

	if saldo <= 0 || valorSaq > saldo {
		fmt.Println("Saldo insuficiente")
		fmt.Printf("Seu saldo é R$ %.2f\n", saldo)
	} else if valorSaq > 500 {
		fmt.Println("Saque máximo permitido de R$ 500,00")
		fmt.Printf("Seu saldo é R$ %.2f\n", saldo)
	} else {
		if numero_saques >= 3 {
			fmt.Println("Limite de saques diários excedido")
			return
		}
		saldo -= valorSaq
		operacoes = append(operacoes, Operacao{Tipo: "Saque", Valor: valorSaq})
		numero_saques++
		fmt.Printf("Saque realizado no valor de R$ %.2f\n", valorSaq)
		fmt.Printf("Seu saldo é R$ %.2f\n", saldo)
	}

}

func extrato() {
	fmt.Println("EXTRATO")
	for i, op := range operacoes {
		fmt.Printf("%d - ", i+1)
		if op.Tipo == "Depósito" {
			fmt.Printf("Depósito no valor de R$ %.2f\n", op.Valor)
		} else if op.Tipo == "Saque" {
			fmt.Printf("Saque no valor de R$ %.2f\n", op.Valor)
		}
	}
	fmt.Printf("\nSaldo final R$ %v\n", saldo)
}

func main() {
	var opcao string

	for {
		fmt.Println("Sistema Bancário")
		fmt.Println("[d] Depósito")
		fmt.Println("[s] Saque")
		fmt.Println("[e] Extrato")
		fmt.Println("[q] Sair")
		fmt.Print("Digite a opção desejada: ")
		fmt.Scanln(&opcao)

		if opcao == "q" {
			fmt.Println("Saindo do sistema.")
			break // Sai do loop for e encerra o programa
		} else if opcao == "d" {
			depositar()
		} else if opcao == "s" {
			sacar()
		} else if opcao == "e" {
			extrato()
		} else {
			fmt.Println("Opção inválida")
		}
	}

}
