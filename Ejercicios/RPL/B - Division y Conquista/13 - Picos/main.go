// Se tiene un arreglo de N >= 3 elementos en forma de pico, esto es: estrictamente
// creciente hasta una determinada posición p, y estrictamente decreciente a partir
// de ella (con 0 < p < N - 1). Por ejemplo, en el arreglo [1, 2, 3, 1, 0, -2] la
// posición del pico es p = 2. Se pide:

// Implementar un algoritmo de división y conquista de orden O(log n) que encuentre
// la posición p del pico: func PosicionPico(v []int, ini, fin int) int.
// La función será invocada inicialmente como: PosicionPico(v, 0, len(v)-1),
// y tiene como pre-condición que el arreglo tenga forma de pico.

// Justificar el orden del algoritmo mediante el teorema maestro.

package main

import "fmt"

// A = 1, B = 2, C = 0
// log en base 2 de 1 = 0, == C entonces O(n^c * log(n)) -> O(log(n))
func PosicionPico(array []int, inicio, fin int) int {
	if inicio == fin {
		return inicio
	}
	mitad := (inicio + fin) / 2
	if array[mitad] > array[mitad+1] {
		return PosicionPico(array, inicio, mitad)
	}
	return PosicionPico(array, mitad+1, fin)
}

func main() {
	arregloPico := []int{1, 2, 3, 1, 0, -2}
	fmt.Println(PosicionPico(arregloPico, 0, len(arregloPico)-1))
}
