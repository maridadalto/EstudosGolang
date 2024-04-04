package main

import (
    "fmt" 
)

func calcularMedia(nota1, nota2 float64) float64 { 
    return (nota1 + nota2) / 2 // Calcula e retorna a média das duas notas
}

func main() { 
    var nomeAluno string 
    var nota1, nota2 float64 

    // Solicita ao usuário que insira o nome do aluno
    fmt.Print("Digite o nome do aluno: ")
	//Lê o input do console, uso o & para obter o endereço da memória dessa variável (ponteiro)
    fmt.Scanln(&nomeAluno)

    // Solicita ao usuário que insira as duas notas
    fmt.Print("Digite a primeira nota: ")
    fmt.Scanln(&nota1)

    fmt.Print("Digite a segunda nota: ")
    fmt.Scanln(&nota2)

    // Chama a função calcularMedia, passando as duas notas como argumentos, e armazena o resultado na variável media
    media := calcularMedia(nota1, nota2)

    // Imprime na tela o nome do aluno, suas duas notas e a média calculada %s e %f são usados para indicar o formato do dado que será exibido %s(string) e %f(float)
	// O .2 é para indicar número de casas decimais após a vírgula
    fmt.Printf("O aluno %s obteve as notas %.2f e %.2f, e a média final é: %.2f\n", nomeAluno, nota1, nota2, media)
}
