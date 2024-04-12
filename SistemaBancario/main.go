package main

import (
	"fmt"
)

var depositos []float64
var saques []float64
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
		depositos = append(depositos, valorDep)
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
		saques = append(saques, valorSaq)
		numero_saques++
		fmt.Printf("Saque realizado no valor de R$ %.2f\n", valorSaq)
		fmt.Printf("Seu saldo é R$ %.2f\n", saldo)
	}

}

func extrato() {
	fmt.Println("EXTRATO")

	for _, dep := range depositos {
		fmt.Printf("Depósito no valor de %.2f\n", dep)
	}

	for _, saq := range saques {
		fmt.Printf("Saque no valor de %.2f\n", saq)
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
