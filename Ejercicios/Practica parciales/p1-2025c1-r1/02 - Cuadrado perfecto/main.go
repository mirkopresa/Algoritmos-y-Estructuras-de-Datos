// Implementar una función func esCuadradoPerfecto(n int) bool que por División y Conquista determine si el
// número n (un positivo entero) es un cuadrado perfecto. Un número es cuadrado perfecto si existe un número entero x
// tal que x²=n.
// Indicar y justificar la complejidad del algoritmo utilizando el Teorema Maestro.

package main

import "fmt"

func cuadradoRecursivo(n, min, max int) bool {
	if min > max {
		return false
	}
	mitad := (min + max) / 2
	if mitad*mitad == n {
		return true
	} else if mitad*mitad > n {
		return cuadradoRecursivo(n, min, mitad-1)
	} else {
		return cuadradoRecursivo(n, mitad+1, max)
	}
}

func esCuadradoPerfecto(n int) bool {
	return cuadradoRecursivo(n, 1, n)
}

func main() {
	fmt.Println(esCuadradoPerfecto(25))
}
