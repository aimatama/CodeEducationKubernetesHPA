package main

import "testing"

func TestRaizQuadrada(t *testing.T) {

	var numero float64 = 0
	var valorOK float64 = 0

	numero = 9
	valorOK = 3

	got := raizQuadrada(numero)
	if got != 3 {
		t.Fatalf("O resultado esperado deveria ter sido '%v' porém o retorno foi igual a '%v'.", valorOK, got)
	}
}
