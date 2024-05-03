package main

import "testing"

func ShouldSumCorrect(t *testing.T) {
	teste := Soma(10, 5)
	resultado := 15
	if teste != resultado {
		t.Error("Valor esperado", resultado, "Valor retornado", teste)
	}
}
func ShouldSumIncorrect(t *testing.T) {
	teste := Soma(10, 5)
	resultado := 10
	if teste != resultado {
		t.Error("Valor esperado", resultado, "Valor retornado", teste)
	}
}
func ShouldSubCorrect(t *testing.T) {
	teste := Subtracao(10, 5)
	resultado := 5
	if teste != resultado {
		t.Error("Valor esperado", resultado, "Valor retornado", teste)
	}
}
func ShouldSubIncorrect(t *testing.T) {
	teste := Subtracao(10, 5)
	resultado := 4
	if teste != resultado {
		t.Error("Valor esperado", resultado, "Valor retornado", teste)
	}
}
func ShouldMultCorrect(t *testing.T) {
	teste := Multiplica(10, 5)
	resultado := 50
	if teste != resultado {
		t.Error("Valor esperado", resultado, "Valor retornado", teste)
	}
}
func ShouldMultIncorrect(t *testing.T) {
	teste := Multiplica(10, 5)
	resultado := 5
	if teste != resultado {
		t.Error("Valor esperado", resultado, "Valor retornado", teste)
	}
}
func ShouldDivCorrect(t *testing.T) {
	teste := Divisao(10, 5)
	resultado := 2
	if teste != resultado {
		t.Error("Valor esperado", resultado, "Valor retornado", teste)
	}
}
func ShouldDivIncorrect(t *testing.T) {
	teste := Divisao(10, 5)
	resultado := 5
	if teste != resultado {
		t.Error("Valor esperado", resultado, "Valor retornado", teste)
	}
}
