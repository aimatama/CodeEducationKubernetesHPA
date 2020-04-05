package main

import (
	"fmt"
	"math"
)

func raizQuadrada(numero float64) float64 {
	return math.Sqrt(numero)
}

func main() {

	var numero float64 = 9
	var resultado float64 = 0
	var limite int = 10

	for j := 1; j <= limite; j++ {
		resultado = raizQuadrada(numero)
		fmt.Println(resultado)
	}

	fmt.Println("Code.Education Rocks!")

}
